# mount

检查某文件系统的挂载情况，通过 `mount.toml` 配置文件指定需要检查的文件系统。

通过`mount`命令实现。该功能实际上也可使用`inputs.exec`插件实现。但考虑到作为业务常用功能，为保证语义明确，特独立出来。