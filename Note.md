# 流媒体播放器的开发笔记
[TOC]

## 设计的方式
> 项目采用的是前后端分离的设计思想,为什么要采用前后端分离的方式尼?主要有以下几点:
> * 前后端解耦是目前非常流行的web网站架构
> * 前端页面和服务通过普通的web引擎渲染
> * 后端的数据是通过渲染后的页面的嵌入脚本调用后端的数据处理和呈现的
> * 解放成产力,提高合作效率(和以前的区别较大,后端不需要了解前端的知识,完全来集中搞后端)
> * 松耦合的架构更加的灵活,部署也是非常的方便,更加的符合微服务的设计特征
> * 有助于项目整体的性能上的一个提升,可靠性的提升
## API设计
### 用户的API设计

* 创建(注册用户):
    - URL:/user
    - Method:POST
    - SC:201,400,500
* 用户登录:
    - URL:/user/:username
    - Method:POST
    - SC:200,400,500
* 获取用户的基本信息:
    - URL:/user/:username
    - Method:GET
    - SC:200,400,401,403,500
* 用户注销：
    - URL:/user/:username
    - Method:DELETE
    - SC:204,400,401,403,500

### 用户的资源API设计
* List all videos:
    - URL:/user/:username/videos
    - Method:GET
    - SC:200,400,500
* Get one video:
    - URL:/user/:username/videus/:vid-id
    - Method:GET
    - SC:200,400,500
* Delete one video:
    - URL:/user/:username/videos/:vid-id
    - Method:DELETE
    - SC:204,400,401,403,500

### 评论的API设计
* Show comments：
    - URL:/videos/:vid-id/comments
    - Method:GET
    - SC:200,400,500
* Post a comment:
    - URL: /videos/:vid-id/comments
    - Method:POST
    - SC:201,400,500
* Delete a comment:
    - URL:/videos/:vid-id/comment/:comment-id
    - Method:DELETE
    - SC:204,400,401,403,500
