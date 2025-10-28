window.addEventListener("DOMContentLoaded", () => {
  const ensureEthers = () => {
    if (typeof window.ethers === "undefined") {
      throw new Error("ethers.js 未加载，请检查网络或刷新页面。");
    }
    return window.ethers;
  };

  const $ = (sel) => document.querySelector(sel);
  const $$ = (sel) => document.querySelectorAll(sel);
  const logPane = $("#log-pane");

  const state = {
    provider: null,
    signer: null,
    account: null,
    stakeContract: null,
    stakeAddress: "",
    rewardToken: "",
    pools: [],
  };

  const stakeAbi = [
    "function startBlock() view returns (uint256)",
    "function endBlock() view returns (uint256)",
    "function fukuaPerBlock() view returns (uint256)",
    "function totalPoolWeight() view returns (uint256)",
    "function claimPaused() view returns (bool)",
    "function withdrawPaused() view returns (bool)",
    "function paused() view returns (bool)",
    "function fukuaToken() view returns (address)",
    "function poolLength() view returns (uint256)",
    "function pool(uint256 pid) view returns (address,uint256,uint256,uint256,uint256,uint256,uint256)",
    "function stakingBalance(uint256,address) view returns (uint256)",
    "function pendingFukua(uint256,address) view returns (uint256)",
    "function withdrawAmount(uint256,address) view returns (uint256,uint256)",
    "function depositEth() payable",
    "function deposit(uint256,uint256)",
    "function unstake(uint256,uint256)",
    "function withdraw(uint256)",
    "function claim(uint256)",
    "function pause()",
    "function unpause()",
    "function pauseWithdraw()",
    "function unpauseWithdraw()",
    "function pauseClaim()",
    "function unpauseClaim()",
    "function setFukuaPerBlock(uint256)",
    "function setPoolWeight(uint256,uint256,bool)",
    "function ADMIN_ROLE() view returns (bytes32)",
    "function UPGRADE_ROLE() view returns (bytes32)",
    "function grantRole(bytes32,address)",
    "function hasRole(bytes32,address) view returns (bool)",
  ];

  const erc20Abi = [
    "function approve(address spender, uint256 amount) returns (bool)",
    "function allowance(address owner, address spender) view returns (uint256)",
    "function decimals() view returns (uint8)",
  ];

  function log(message, type = "info") {
    const timestamp = new Date().toLocaleTimeString();
    if (logPane) {
      logPane.textContent = `[${timestamp}] [${type.toUpperCase()}] ${message}\n${logPane.textContent}`;
    } else {
      console.log(message);
    }
  }

  async function connectWallet() {
    let ethersLib;
    try {
      ethersLib = ensureEthers();
    } catch (err) {
      log(err.message, "error");
      return;
    }

    if (!window.ethereum) {
      log("请安装 MetaMask 或兼容钱包扩展。", "error");
      return;
    }

    try {
      log("正在请求钱包授权...");
      await window.ethereum.request({ method: "eth_requestAccounts" });
      state.provider = new ethersLib.BrowserProvider(window.ethereum);
      state.signer = await state.provider.getSigner();
      state.account = await state.signer.getAddress();

      $("#wallet-address").textContent = state.account;
      log(`已连接钱包：${state.account}`);
      await refreshUserState();
    } catch (err) {
      log(`连接钱包失败：${err.message ?? err}`, "error");
    }
  }

  async function loadContract() {
    if (!state.signer) {
      log("请先连接钱包。", "warn");
      return;
    }

    let ethersLib;
    try {
      ethersLib = ensureEthers();
    } catch (err) {
      log(err.message, "error");
      return;
    }

    const stakeInputEl = $("#stake-address");
    const addrInput = (stakeInputEl?.value || "").trim();

    let normalizedAddress;
    try {
      normalizedAddress = ethersLib.getAddress(addrInput);
    } catch (err) {
      try {
        normalizedAddress = ethersLib.getAddress(addrInput.toLowerCase());
      } catch (err2) {
        log(
          `请输入合法的 FukuaStake 合约地址。${err2?.message ? " 详情: " + err2.message : ""}`,
          "warn"
        );
        return;
      }
    }

    state.stakeAddress = normalizedAddress;
    state.stakeContract = new ethersLib.Contract(state.stakeAddress, stakeAbi, state.signer);

    try {
      await refreshAllState();
      $("#contract-status").textContent = "合约加载成功";
      $("#global-state").hidden = false;
      $("#pool-section").hidden = false;
      $("#user-section").hidden = false;
      $("#admin-section").hidden = false;
      log(`已连接 FukuaStake 合约：${state.stakeAddress}`);
    } catch (err) {
      log(`加载合约失败：${err.message ?? err}`, "error");
    }
  }

  async function refreshAllState() {
    if (!state.stakeContract) return;
    const ethersLib = ensureEthers();

    const [
      startBlock,
      endBlock,
      fukuaPerBlock,
      totalWeight,
      claimPaused,
      withdrawPaused,
      paused,
      rewardToken,
      poolLength,
    ] = await Promise.all([
      state.stakeContract.startBlock(),
      state.stakeContract.endBlock(),
      state.stakeContract.fukuaPerBlock(),
      state.stakeContract.totalPoolWeight(),
      state.stakeContract.claimPaused(),
      state.stakeContract.withdrawPaused(),
      state.stakeContract.paused(),
      state.stakeContract.fukuaToken(),
      state.stakeContract.poolLength(),
    ]);

    state.rewardToken = rewardToken;

    $("#start-block").textContent = startBlock.toString();
    $("#end-block").textContent = endBlock.toString();
    $("#reward-per-block").textContent = ethersLib.formatEther(fukuaPerBlock);
    $("#total-weight").textContent = totalWeight.toString();
    $("#paused-flag").textContent = paused ? "是" : "否";
    $("#claim-paused-flag").textContent = claimPaused ? "是" : "否";
    $("#withdraw-paused-flag").textContent = withdrawPaused ? "是" : "否";
    $("#reward-token").textContent = rewardToken;
    $("#reward-address").value = rewardToken;

    state.pools = [];
    const tableBody = $("#pool-table-body");
    tableBody.innerHTML = "";
    const pidSelect = $("#user-pid");
    pidSelect.innerHTML = "";

    for (let pid = 0; pid < Number(poolLength); pid++) {
      const pool = await state.stakeContract.pool(pid);
      state.pools.push(pool);

      const row = document.createElement("tr");
      row.innerHTML = `
        <td>${pid}</td>
        <td>${pool[0]}</td>
        <td>${pool[1]}</td>
        <td>${ethersLib.formatEther(pool[4])}</td>
        <td>${ethersLib.formatEther(pool[5])}</td>
        <td>${pool[6]}</td>
      `;
      tableBody.appendChild(row);

      const option = document.createElement("option");
      option.value = pid;
      option.textContent = `PID ${pid}`;
      pidSelect.appendChild(option);
    }

    if (state.account) {
      await refreshUserState();
    }
  }

  async function refreshUserState() {
    if (!state.stakeContract || !state.account) return;
    const ethersLib = ensureEthers();
    const pid = Number($("#user-pid").value || 0);

    const [stakeBal, pending, withdrawInfo] = await Promise.all([
      state.stakeContract.stakingBalance(pid, state.account),
      state.stakeContract.pendingFukua(pid, state.account),
      state.stakeContract.withdrawAmount(pid, state.account),
    ]);

    $("#user-stake").textContent = ethersLib.formatEther(stakeBal);
    $("#user-pending").textContent = ethersLib.formatEther(pending);
    $("#user-requested").textContent = ethersLib.formatEther(withdrawInfo[0]);
    $("#user-unlocked").textContent = ethersLib.formatEther(withdrawInfo[1]);
  }

  async function approveErc20() {
    if (!state.stakeContract) {
      log("请先加载合约。", "warn");
      return;
    }
    try {
      const ethersLib = ensureEthers();
      const tokenAddr = $("#stake-token").value.trim();
      if (!ethersLib.isAddress(tokenAddr)) {
        log("请输入合法的 ERC20 代币地址。", "warn");
        return;
      }
      const amount = $("#stake-amount").value.trim();
      if (!amount) {
        log("请输入授权额度。", "warn");
        return;
      }
      const token = new ethersLib.Contract(tokenAddr, erc20Abi, state.signer);
      const decimals = await token.decimals();
      const value = ethersLib.parseUnits(amount, decimals);
      const tx = await token.approve(state.stakeAddress, value);
      log(`发送 approve 交易：${tx.hash}`);
      await tx.wait();
      log("授权成功。");
    } catch (err) {
      log(`授权失败：${err.message ?? err}`, "error");
    }
  }

  async function deposit() {
    if (!state.stakeContract || !state.pools.length) {
      log("请先加载合约并选择质押池。", "warn");
      return;
    }
    const pid = Number($("#user-pid").value || 0);
    const amount = $("#stake-amount").value.trim();
    if (!amount) {
      log("请输入质押数量。", "warn");
      return;
    }

    const pool = state.pools[pid];
    if (!pool) {
      log("未找到对应池信息。", "warn");
      return;
    }

    try {
      const ethersLib = ensureEthers();
      if (pool[0] === ethersLib.ZeroAddress) {
        const value = ethersLib.parseEther(amount);
        const tx = await state.stakeContract.depositEth({ value });
        log(`发送 depositEth 交易：${tx.hash}`);
        await tx.wait();
        log("质押成功。");
      } else {
        const tokenAddr = $("#stake-token").value.trim();
        if (!ethersLib.isAddress(tokenAddr)) {
          log("该池需要 ERC20 质押，请填写代币地址。", "warn");
          return;
        }
        const token = new ethersLib.Contract(tokenAddr, erc20Abi, state.signer);
        const decimals = await token.decimals();
        const allowance = await token.allowance(state.account, state.stakeAddress);
        const value = ethersLib.parseUnits(amount, decimals);
        if (allowance < value) {
          const txApprove = await token.approve(state.stakeAddress, value);
          log(`Allowance 不足，已发送 approve 交易：${txApprove.hash}`);
          await txApprove.wait();
        }
        const tx = await state.stakeContract.deposit(pid, value);
        log(`发送 deposit 交易：${tx.hash}`);
        await tx.wait();
        log("质押成功。");
      }
      await refreshAllState();
    } catch (err) {
      log(`质押失败：${err.message ?? err}`, "error");
    }
  }

  async function claim() {
    if (!state.stakeContract) {
      log("请先加载合约。", "warn");
      return;
    }
    const pid = Number($("#user-pid").value || 0);
    try {
      const tx = await state.stakeContract.claim(pid);
      log(`发送 claim 交易：${tx.hash}`);
      await tx.wait();
      log("领取成功。");
      await refreshAllState();
    } catch (err) {
      log(`领取失败：${err.message ?? err}`, "error");
    }
  }

  async function unstake() {
    if (!state.stakeContract) {
      log("请先加载合约。", "warn");
      return;
    }
    const pid = Number($("#user-pid").value || 0);
    const amount = $("#unstake-amount").value.trim();
    if (!amount) {
      log("请输入解质押数量。", "warn");
      return;
    }

    try {
      const ethersLib = ensureEthers();
      const value = ethersLib.parseEther(amount);
      const tx = await state.stakeContract.unstake(pid, value);
      log(`发送 unstake 交易：${tx.hash}`);
      await tx.wait();
      log("解质押请求成功。");
      await refreshAllState();
    } catch (err) {
      log(`解质押失败：${err.message ?? err}`, "error");
    }
  }

  async function withdraw() {
    if (!state.stakeContract) {
      log("请先加载合约。", "warn");
      return;
    }
    const pid = Number($("#user-pid").value || 0);
    try {
      const tx = await state.stakeContract.withdraw(pid);
      log(`发送 withdraw 交易：${tx.hash}`);
      await tx.wait();
      log("提现成功。");
      await refreshAllState();
    } catch (err) {
      log(`提现失败：${err.message ?? err}`, "error");
    }
  }

  function makeAdminHandler(methodName, message) {
    return async () => {
      try {
        const tx = await state.stakeContract[methodName]();
        log(`${message} 交易发送：${tx.hash}`);
        await tx.wait();
        log(`${message} 完成。`);
        await refreshAllState();
      } catch (err) {
        log(`${message} 失败：${err.message ?? err}`, "error");
      }
    };
  }

  async function updateRewardRate() {
    if (!state.stakeContract) {
      log("请先加载合约并确认拥有管理员权限。", "warn");
      return;
    }
    const value = $("#new-reward-rate").value.trim();
    if (!value) {
      log("请输入新的 fukuaPerBlock。", "warn");
      return;
    }

    try {
      const ethersLib = ensureEthers();
      const parsed = ethersLib.parseEther(value);
      const tx = await state.stakeContract.setFukuaPerBlock(parsed);
      log(`更新奖励速率交易：${tx.hash}`);
      await tx.wait();
      log("奖励速率已更新。");
      await refreshAllState();
    } catch (err) {
      log(`更新奖励速率失败：${err.message ?? err}`, "error");
    }
  }

  async function updatePoolWeight() {
    if (!state.stakeContract) {
      log("请先加载合约并确认拥有管理员权限。", "warn");
      return;
    }
    const pid = Number($("#pool-weight-pid").value);
    const weight = Number($("#pool-weight-value").value);
    const withUpdate = $("#pool-weight-update-all").checked;
    if (Number.isNaN(pid) || Number.isNaN(weight) || weight <= 0) {
      log("请输入合法的 PID 和权重。", "warn");
      return;
    }
    try {
      const tx = await state.stakeContract.setPoolWeight(pid, weight, withUpdate);
      log(`更新权重交易：${tx.hash}`);
      await tx.wait();
      log("权重已更新。");
      await refreshAllState();
    } catch (err) {
      log(`更新权重失败：${err.message ?? err}`, "error");
    }
  }

  $("#connect-wallet").addEventListener("click", connectWallet);
  $("#load-contract").addEventListener("click", loadContract);
  $("#user-pid").addEventListener("change", refreshUserState);
  $("#approve-button").addEventListener("click", approveErc20);
  $("#deposit-button").addEventListener("click", deposit);
  $("#claim-button").addEventListener("click", claim);
  $("#unstake-button").addEventListener("click", unstake);
  $("#withdraw-button").addEventListener("click", withdraw);
  $("#update-rate-button").addEventListener("click", updateRewardRate);
  $("#update-weight-button").addEventListener("click", updatePoolWeight);

  $("#pause-global").addEventListener("click", makeAdminHandler("pause", "全局暂停"));
  $("#unpause-global").addEventListener("click", makeAdminHandler("unpause", "取消全局暂停"));
  $("#pause-claim").addEventListener("click", makeAdminHandler("pauseClaim", "暂停奖励领取"));
  $("#unpause-claim").addEventListener("click", makeAdminHandler("unpauseClaim", "恢复奖励领取"));
  $("#pause-withdraw").addEventListener("click", makeAdminHandler("pauseWithdraw", "暂停提现"));
  $("#unpause-withdraw").addEventListener("click", makeAdminHandler("unpauseWithdraw", "恢复提现"));

  window.ethereum?.on("accountsChanged", () => {
    connectWallet().then(refreshAllState).catch((err) => {
      if (err) log(`刷新失败：${err.message ?? err}`, "error");
    });
  });
});
