# 项目 Git 提交规范

```text
<emoji><type> [scope]: <subject>

[body]

[footer]
```
- type: 在添加后面 ! 表示这次提交是 Breaking Change 不兼容修改 
  - feat: 新功能
  - fix: 修复问题
  - docs: 修改文档
  - style: 修改代码格式，不影响代码逻辑
  - refactor: 重构代码，理论上不影响现有功能
  - perf: 提升性能
  - test: 增加修改测试用例
  - chore: 修改工具相关（包括但不限于文档、代码生成等）
  - deps: 升级依赖
- scope: 修改文件的范围(docs、pkg、...)
- subject: 这次提交做什么什么
- body: 详细说明
- footer

emoji:

| emoji |          	emoji代码          |  中文   |    	commit说明    |
|:-----:|:--------------------------:|:-----:|:---------------:|
|  🎨   |          	:art:	           |  调色板  |   改进代码结构/代码格式   |
|   ⚡   |           :zap:            |  闪电   |      提升性能       |
|  🐎   |        :racehorse:         |  赛马   |      提升性能       |
|  🔥   |           :fire:           |  火焰   |     移除代码或文件     |
|  🐛   |           :bug:            |  bug  |     修复 bug      |
|  🚑   |        :ambulance:         |  急救车  |      重要补丁       |
|   ✨   |         :sparkles:         |  火花   |      引入新功能      |
|  📝   |          :pencil:          |  铅笔   |      撰写文档       |
|  🚀   |          :rocket:          |  火箭   |      部署功能       |
|  💄   |         :lipstick:         |  口红   |   更新 UI 和样式文件   |
|  🎉   |           :tada:           |  庆祝   |      初次提交       |
|   ✅   |     :white_check_mark:     | 白色复选框 |      增加测试       |
|  🔒   |           :lock:           |   锁   |     修复安全问题      |
|  🍎   |          :apple:           |  苹果   |  修复 macOS 下的问题  |
|  🐧   |         :penguin:          |  企鹅   |  修复 Linux 下的问题  |
|  🏁   |       :checked_flag:       |  旗帜   | 修复 Windows 下的问题 |
|  🔖   |         :bookmark:         |  书签   |     发行/版本标签     |
|  🚨   |      :rotating_light:      |  警车灯  |  移除 linter 警告   |
|  🚧   |       :construction:       |  施工   |      工作进行中      |
|  💚   |       :green_heart:        |  绿心   |   修复 CI 构建问题    |
|   ⬇   |        :arrow_down:        | 下降箭头  |      降级依赖       |
|   ⬆   |         :arrow_up:         | 上升箭头  |      升级依赖       |
|  👷   |   :construction_worker:    |  工人   |   添加 CI 构建系统    |
|  📈   | :chart_with_upwards_trend: | 上升趋势图 |    添加分析或跟踪代码    |
|  🔨   |          :hammer:          |  锤子   |      重大重构       |
|   ➖   |     :heavy_minus_sign:     |  减号   |     减少一个依赖      |
|  🐳   |          :whale:           |  鲸鱼   |      相关工作       |
|   ➕   |     :heavy_plus_sign:      |  加号   |     增加一个依赖      |
|  🔧   |          :wrench:          |  扳手   |     修改配置文件      |
|  🌐   |   :globe_with_meridians:   |  地球   |     国际化与本地化     |
|   ✏   |         :pencil2:          |  铅笔   |     修复 typo     |
