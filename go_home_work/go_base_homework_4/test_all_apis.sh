#!/bin/bash

# 个人博客后端系统 - 自动化API测试脚本
# 使用方法: chmod +x test_all_apis.sh && ./test_all_apis.sh

set -e  # 遇到错误即停止

BASE_URL="http://localhost:8080"
TEST_USER="testuser_$(date +%s)"  # 使用时间戳避免用户名冲突
TEST_PASSWORD="password123"
TEST_EMAIL="test_$(date +%s)@example.com"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 测试计数器
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# 打印函数
print_header() {
    echo -e "\n${BLUE}===========================================${NC}"
    echo -e "${BLUE} $1 ${NC}"
    echo -e "${BLUE}===========================================${NC}\n"
}

print_test() {
    echo -e "${YELLOW}[测试] $1${NC}"
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
}

print_success() {
    echo -e "${GREEN}✓ 通过: $1${NC}"
    PASSED_TESTS=$((PASSED_TESTS + 1))
}

print_error() {
    echo -e "${RED}✗ 失败: $1${NC}"
    FAILED_TESTS=$((FAILED_TESTS + 1))
}

print_info() {
    echo -e "${BLUE}ℹ 信息: $1${NC}"
}

# HTTP请求函数
make_request() {
    local method=$1
    local url=$2
    local data=$3
    local headers=$4
    
    if [[ -n "$headers" ]]; then
        if [[ -n "$data" ]]; then
            curl -s -X "$method" "$url" -H "Content-Type: application/json" -H "$headers" -d "$data"
        else
            curl -s -X "$method" "$url" -H "$headers"
        fi
    else
        if [[ -n "$data" ]]; then
            curl -s -X "$method" "$url" -H "Content-Type: application/json" -d "$data"
        else
            curl -s -X "$method" "$url"
        fi
    fi
}

# 提取JSON字段值
extract_json_field() {
    local json=$1
    local field=$2
    echo "$json" | grep -o "\"$field\":[^,}]*" | sed 's/.*://; s/"//g'
}

# 检查服务器是否运行
check_server() {
    print_header "检查服务器状态"
    print_test "服务器健康检查"
    
    response=$(curl -s "$BASE_URL/health" || echo "ERROR")
    
    if echo "$response" | grep -q "博客后端系统运行正常"; then
        print_success "服务器运行正常"
        return 0
    else
        print_error "服务器未运行或响应异常"
        echo "响应: $response"
        exit 1
    fi
}

# 用户认证测试
test_auth() {
    print_header "用户认证测试"
    
    # 测试用户注册
    print_test "用户注册 - 正常情况"
    register_data="{\"username\":\"$TEST_USER\",\"password\":\"$TEST_PASSWORD\",\"email\":\"$TEST_EMAIL\"}"
    register_response=$(make_request "POST" "$BASE_URL/api/auth/register" "$register_data")
    
    if echo "$register_response" | grep -q "用户注册成功"; then
        print_success "用户注册成功"
        USER_ID=$(extract_json_field "$register_response" "user_id")
        print_info "用户ID: $USER_ID"
    else
        print_error "用户注册失败: $register_response"
    fi
    
    # 测试重复注册
    print_test "用户注册 - 用户名重复"
    duplicate_response=$(make_request "POST" "$BASE_URL/api/auth/register" "$register_data")
    
    if echo "$duplicate_response" | grep -q "用户名已存在"; then
        print_success "正确处理重复用户名"
    else
        print_error "重复用户名处理失败: $duplicate_response"
    fi
    
    # 测试用户登录
    print_test "用户登录 - 正确密码"
    login_data="{\"username\":\"$TEST_USER\",\"password\":\"$TEST_PASSWORD\"}"
    login_response=$(make_request "POST" "$BASE_URL/api/auth/login" "$login_data")
    
    if echo "$login_response" | grep -q "登录成功"; then
        print_success "用户登录成功"
        JWT_TOKEN=$(extract_json_field "$login_response" "token")
        print_info "JWT Token已获取"
    else
        print_error "用户登录失败: $login_response"
        exit 1
    fi
    
    # 测试错误密码登录
    print_test "用户登录 - 错误密码"
    wrong_login_data="{\"username\":\"$TEST_USER\",\"password\":\"wrongpassword\"}"
    wrong_login_response=$(make_request "POST" "$BASE_URL/api/auth/login" "$wrong_login_data")
    
    if echo "$wrong_login_response" | grep -q "用户名或密码错误"; then
        print_success "正确处理错误密码"
    else
        print_error "错误密码处理失败: $wrong_login_response"
    fi
    
    # 测试JWT认证
    print_test "JWT认证验证"
    profile_response=$(make_request "GET" "$BASE_URL/api/protected/profile" "" "Authorization: Bearer $JWT_TOKEN")
    
    if echo "$profile_response" | grep -q "访问成功"; then
        print_success "JWT认证正常工作"
    else
        print_error "JWT认证失败: $profile_response"
    fi
}

# 文章管理测试
test_posts() {
    print_header "文章管理测试"
    
    # 测试获取文章列表（空列表或已有文章）
    print_test "获取文章列表"
    posts_response=$(make_request "GET" "$BASE_URL/api/posts")
    
    if echo "$posts_response" | grep -q "获取文章列表成功"; then
        print_success "文章列表获取成功"
    else
        print_error "文章列表获取失败: $posts_response"
    fi
    
    # 测试创建文章
    print_test "创建文章"
    create_post_data="{\"title\":\"自动化测试文章\",\"content\":\"这是通过自动化测试脚本创建的文章，用于验证API功能正常工作。\"}"
    create_post_response=$(make_request "POST" "$BASE_URL/api/protected/posts" "$create_post_data" "Authorization: Bearer $JWT_TOKEN")
    
    if echo "$create_post_response" | grep -q "自动化测试文章"; then
        print_success "文章创建成功"
        POST_ID=$(extract_json_field "$create_post_response" "id")
        print_info "文章ID: $POST_ID"
    else
        print_error "文章创建失败: $create_post_response"
    fi
    
    # 测试无认证创建文章
    print_test "创建文章 - 无认证"
    unauth_create_response=$(make_request "POST" "$BASE_URL/api/protected/posts" "$create_post_data")
    
    if echo "$unauth_create_response" | grep -q "请提供访问令牌"; then
        print_success "正确拒绝无认证创建"
    else
        print_error "无认证创建处理失败: $unauth_create_response"
    fi
    
    # 测试获取文章详情
    print_test "获取文章详情"
    post_detail_response=$(make_request "GET" "$BASE_URL/api/posts/$POST_ID")
    
    if echo "$post_detail_response" | grep -q "获取文章详情成功"; then
        print_success "文章详情获取成功"
    else
        print_error "文章详情获取失败: $post_detail_response"
    fi
    
    # 测试获取不存在文章
    print_test "获取不存在文章"
    nonexistent_response=$(make_request "GET" "$BASE_URL/api/posts/999999")
    
    if echo "$nonexistent_response" | grep -q "文章不存在"; then
        print_success "正确处理不存在文章"
    else
        print_error "不存在文章处理失败: $nonexistent_response"
    fi
    
    # 测试更新文章
    print_test "更新文章 - 作者权限"
    update_post_data="{\"title\":\"更新后的自动化测试文章\",\"content\":\"这是更新后的文章内容，用于测试更新功能。\"}"
    update_post_response=$(make_request "PUT" "$BASE_URL/api/protected/posts/$POST_ID" "$update_post_data" "Authorization: Bearer $JWT_TOKEN")
    
    if echo "$update_post_response" | grep -q "更新后的自动化测试文章"; then
        print_success "文章更新成功"
    else
        print_error "文章更新失败: $update_post_response"
    fi
    
    # 测试更新不存在文章
    print_test "更新不存在文章"
    update_nonexistent_response=$(make_request "PUT" "$BASE_URL/api/protected/posts/999999" "$update_post_data" "Authorization: Bearer $JWT_TOKEN")
    
    if echo "$update_nonexistent_response" | grep -q "文章不存在"; then
        print_success "正确处理更新不存在文章"
    else
        print_error "更新不存在文章处理失败: $update_nonexistent_response"
    fi
}

# 评论管理测试
test_comments() {
    print_header "评论管理测试"
    
    # 测试创建评论
    print_test "创建评论"
    create_comment_data="{\"content\":\"这是通过自动化测试脚本添加的评论，用于验证评论功能。\"}"
    create_comment_response=$(make_request "POST" "$BASE_URL/api/protected/posts/$POST_ID/comments" "$create_comment_data" "Authorization: Bearer $JWT_TOKEN")
    
    if echo "$create_comment_response" | grep -q "通过自动化测试脚本添加的评论"; then
        print_success "评论创建成功"
        COMMENT_ID=$(extract_json_field "$create_comment_response" "id")
        print_info "评论ID: $COMMENT_ID"
    else
        print_error "评论创建失败: $create_comment_response"
    fi
    
    # 测试无认证创建评论
    print_test "创建评论 - 无认证"
    unauth_comment_response=$(make_request "POST" "$BASE_URL/api/protected/posts/$POST_ID/comments" "$create_comment_data")
    
    if echo "$unauth_comment_response" | grep -q "请提供访问令牌"; then
        print_success "正确拒绝无认证评论"
    else
        print_error "无认证评论处理失败: $unauth_comment_response"
    fi
    
    # 测试对不存在文章评论
    print_test "对不存在文章评论"
    comment_nonexistent_response=$(make_request "POST" "$BASE_URL/api/protected/posts/999999/comments" "$create_comment_data" "Authorization: Bearer $JWT_TOKEN")
    
    if echo "$comment_nonexistent_response" | grep -q "文章不存在"; then
        print_success "正确处理对不存在文章评论"
    else
        print_error "对不存在文章评论处理失败: $comment_nonexistent_response"
    fi
    
    # 测试获取评论列表
    print_test "获取评论列表"
    comments_response=$(make_request "GET" "$BASE_URL/api/posts/$POST_ID/comments")
    
    if echo "$comments_response" | grep -q "获取评论列表成功"; then
        print_success "评论列表获取成功"
        comment_count=$(extract_json_field "$comments_response" "count")
        print_info "评论数量: $comment_count"
    else
        print_error "评论列表获取失败: $comments_response"
    fi
    
    # 测试获取不存在文章的评论
    print_test "获取不存在文章评论列表"
    nonexistent_comments_response=$(make_request "GET" "$BASE_URL/api/posts/999999/comments")
    
    if echo "$nonexistent_comments_response" | grep -q "文章不存在"; then
        print_success "正确处理不存在文章评论列表"
    else
        print_error "不存在文章评论列表处理失败: $nonexistent_comments_response"
    fi
}

# 权限和安全测试
test_security() {
    print_header "权限和安全测试"
    
    # 测试无效JWT
    print_test "无效JWT访问"
    invalid_jwt_response=$(make_request "GET" "$BASE_URL/api/protected/profile" "" "Authorization: Bearer invalid_token")
    
    if echo "$invalid_jwt_response" | grep -q "无效的访问令牌"; then
        print_success "正确处理无效JWT"
    else
        print_error "无效JWT处理失败: $invalid_jwt_response"
    fi
    
    # 测试删除文章（作者权限）
    print_test "删除文章 - 作者权限"
    delete_response=$(make_request "DELETE" "$BASE_URL/api/protected/posts/$POST_ID" "" "Authorization: Bearer $JWT_TOKEN")
    
    if echo "$delete_response" | grep -q "文章删除成功"; then
        print_success "文章删除成功"
    else
        print_error "文章删除失败: $delete_response"
    fi
    
    # 测试删除不存在文章
    print_test "删除不存在文章"
    delete_nonexistent_response=$(make_request "DELETE" "$BASE_URL/api/protected/posts/999999" "" "Authorization: Bearer $JWT_TOKEN")
    
    if echo "$delete_nonexistent_response" | grep -q "文章不存在"; then
        print_success "正确处理删除不存在文章"
    else
        print_error "删除不存在文章处理失败: $delete_nonexistent_response"
    fi
}

# 清理测试数据
cleanup() {
    print_header "清理测试数据"
    print_info "测试完成，部分测试数据已自动清理"
    print_info "注册的测试用户: $TEST_USER"
}

# 打印测试总结
print_summary() {
    print_header "测试总结"
    echo -e "${BLUE}总测试数: $TOTAL_TESTS${NC}"
    echo -e "${GREEN}通过: $PASSED_TESTS${NC}"
    echo -e "${RED}失败: $FAILED_TESTS${NC}"
    
    success_rate=$((PASSED_TESTS * 100 / TOTAL_TESTS))
    echo -e "${BLUE}成功率: ${success_rate}%${NC}"
    
    if [ $FAILED_TESTS -eq 0 ]; then
        echo -e "\n${GREEN}🎉 所有测试通过！博客系统API功能完全正常！${NC}\n"
        exit 0
    else
        echo -e "\n${RED}❌ 有 $FAILED_TESTS 个测试失败，请检查系统状态${NC}\n"
        exit 1
    fi
}

# 主程序
main() {
    print_header "个人博客后端系统 - 自动化API测试"
    print_info "开始全面测试所有API接口..."
    print_info "测试用户: $TEST_USER"
    print_info "服务器地址: $BASE_URL"
    
    check_server
    test_auth
    test_posts
    test_comments  
    test_security
    cleanup
    print_summary
}

# 检查curl是否安装
if ! command -v curl &> /dev/null; then
    echo -e "${RED}错误: 需要安装curl来运行测试${NC}"
    exit 1
fi

# 运行主程序
main