# Go-utilsLog
这是一个可扩展可备份收集Debug,Info,Error等级别的日志库，存储类型为文件或控制台输出


用法
	logger := ckLog.LogNewFileLogger("./", "log")
	logger.Debug(ckLog.DebugLevel, "测试哈哈")

	defer logger.Close()
  
  性能测试无限 循环
  文件超过1024 * 1024会自动备份，从新创建，可以自行测试或任意修改，欢迎完善项目，star
  
