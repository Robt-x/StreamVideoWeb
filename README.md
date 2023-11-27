# StreamVideoWeb
API设计：
用户：
    创建/注册用户:  URL：/user Method:POST, SC:201(success),400,500
    用户登录:   URL: /user/:username Method:POST,SC:200(ok),400,500
    获取用户基本信息: URL: /user/:username Method:GET,SC:200,400,401(没有验证),403(验证后没有操作资源的某项权限),500
    用户注销: URL:/user/:username Method: DELETE,SC:204,400,401,403,405

用户资源：
    List all videos: URL:/user/:username/videos Method:GET, SC:200,400,500
    Get one video: URL:/user/:username/videos/:vid-id Method:GET.SC:200,400,500
    Delete ont video: URL:/user/:username/videos/:vid-id Method:DELETE, SC:204,400,401,403,500

评论：
    Show comments: URL:/videos/:vid-id/comments Method: GET,SC:200,400,500
    Post a comment: URL:/videos/:vid-id/comments Method: POST,SC:201,400,500
    Delete a comment: URL:/videos/:vid-id/comment/:comment-id Method: DELETE,SC:204,400,401,403,500