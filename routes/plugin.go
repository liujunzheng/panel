package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"github.com/TheTNB/panel/app/http/controllers/plugins"
	"github.com/TheTNB/panel/app/http/middleware"
)

// Plugin 加载插件路由
func Plugin() {
	facades.Route().Prefix("api/plugins").Middleware(middleware.Jwt(), middleware.MustInstall()).Group(func(r route.Router) {
		r.Prefix("openresty").Group(func(route route.Router) {
			openRestyController := plugins.NewOpenrestyController()
			route.Get("status", openRestyController.Status)
			route.Post("reload", openRestyController.Reload)
			route.Post("start", openRestyController.Start)
			route.Post("stop", openRestyController.Stop)
			route.Post("restart", openRestyController.Restart)
			route.Get("load", openRestyController.Load)
			route.Get("config", openRestyController.GetConfig)
			route.Post("config", openRestyController.SaveConfig)
			route.Get("errorLog", openRestyController.ErrorLog)
			route.Post("clearErrorLog", openRestyController.ClearErrorLog)
		})
		r.Prefix("mysql57").Group(func(route route.Router) {
			mySQLController := plugins.NewMySQLController()
			route.Get("status", mySQLController.Status)
			route.Post("reload", mySQLController.Reload)
			route.Post("start", mySQLController.Start)
			route.Post("stop", mySQLController.Stop)
			route.Post("restart", mySQLController.Restart)
			route.Get("load", mySQLController.Load)
			route.Get("config", mySQLController.GetConfig)
			route.Post("config", mySQLController.SaveConfig)
			route.Get("errorLog", mySQLController.ErrorLog)
			route.Post("clearErrorLog", mySQLController.ClearErrorLog)
			route.Get("slowLog", mySQLController.SlowLog)
			route.Post("clearSlowLog", mySQLController.ClearSlowLog)
			route.Get("rootPassword", mySQLController.GetRootPassword)
			route.Post("rootPassword", mySQLController.SetRootPassword)
			route.Get("databases", mySQLController.DatabaseList)
			route.Post("databases", mySQLController.AddDatabase)
			route.Delete("databases", mySQLController.DeleteDatabase)
			route.Get("backups", mySQLController.BackupList)
			route.Post("backups", mySQLController.CreateBackup)
			route.Put("backups", mySQLController.UploadBackup)
			route.Delete("backups", mySQLController.DeleteBackup)
			route.Post("backups/restore", mySQLController.RestoreBackup)
			route.Get("users", mySQLController.UserList)
			route.Post("users", mySQLController.AddUser)
			route.Delete("users", mySQLController.DeleteUser)
			route.Post("users/password", mySQLController.SetUserPassword)
			route.Post("users/privileges", mySQLController.SetUserPrivileges)
		})
		r.Prefix("mysql80").Group(func(route route.Router) {
			mySQLController := plugins.NewMySQLController()
			route.Get("status", mySQLController.Status)
			route.Post("reload", mySQLController.Reload)
			route.Post("start", mySQLController.Start)
			route.Post("stop", mySQLController.Stop)
			route.Post("restart", mySQLController.Restart)
			route.Get("load", mySQLController.Load)
			route.Get("config", mySQLController.GetConfig)
			route.Post("config", mySQLController.SaveConfig)
			route.Get("errorLog", mySQLController.ErrorLog)
			route.Post("clearErrorLog", mySQLController.ClearErrorLog)
			route.Get("slowLog", mySQLController.SlowLog)
			route.Post("clearSlowLog", mySQLController.ClearSlowLog)
			route.Get("rootPassword", mySQLController.GetRootPassword)
			route.Post("rootPassword", mySQLController.SetRootPassword)
			route.Get("databases", mySQLController.DatabaseList)
			route.Post("databases", mySQLController.AddDatabase)
			route.Delete("databases", mySQLController.DeleteDatabase)
			route.Get("backups", mySQLController.BackupList)
			route.Post("backups", mySQLController.CreateBackup)
			route.Put("backups", mySQLController.UploadBackup)
			route.Delete("backups", mySQLController.DeleteBackup)
			route.Post("backups/restore", mySQLController.RestoreBackup)
			route.Get("users", mySQLController.UserList)
			route.Post("users", mySQLController.AddUser)
			route.Delete("users", mySQLController.DeleteUser)
			route.Post("users/password", mySQLController.SetUserPassword)
			route.Post("users/privileges", mySQLController.SetUserPrivileges)
		})
		r.Prefix("mysql84").Group(func(route route.Router) {
			mySQLController := plugins.NewMySQLController()
			route.Get("status", mySQLController.Status)
			route.Post("reload", mySQLController.Reload)
			route.Post("start", mySQLController.Start)
			route.Post("stop", mySQLController.Stop)
			route.Post("restart", mySQLController.Restart)
			route.Get("load", mySQLController.Load)
			route.Get("config", mySQLController.GetConfig)
			route.Post("config", mySQLController.SaveConfig)
			route.Get("errorLog", mySQLController.ErrorLog)
			route.Post("clearErrorLog", mySQLController.ClearErrorLog)
			route.Get("slowLog", mySQLController.SlowLog)
			route.Post("clearSlowLog", mySQLController.ClearSlowLog)
			route.Get("rootPassword", mySQLController.GetRootPassword)
			route.Post("rootPassword", mySQLController.SetRootPassword)
			route.Get("databases", mySQLController.DatabaseList)
			route.Post("databases", mySQLController.AddDatabase)
			route.Delete("databases", mySQLController.DeleteDatabase)
			route.Get("backups", mySQLController.BackupList)
			route.Post("backups", mySQLController.CreateBackup)
			route.Put("backups", mySQLController.UploadBackup)
			route.Delete("backups", mySQLController.DeleteBackup)
			route.Post("backups/restore", mySQLController.RestoreBackup)
			route.Get("users", mySQLController.UserList)
			route.Post("users", mySQLController.AddUser)
			route.Delete("users", mySQLController.DeleteUser)
			route.Post("users/password", mySQLController.SetUserPassword)
			route.Post("users/privileges", mySQLController.SetUserPrivileges)
		})
		r.Prefix("postgresql15").Group(func(route route.Router) {
			postgresql15Controller := plugins.NewPostgresql15Controller()
			route.Get("status", postgresql15Controller.Status)
			route.Post("reload", postgresql15Controller.Reload)
			route.Post("start", postgresql15Controller.Start)
			route.Post("stop", postgresql15Controller.Stop)
			route.Post("restart", postgresql15Controller.Restart)
			route.Get("load", postgresql15Controller.Load)
			route.Get("config", postgresql15Controller.GetConfig)
			route.Post("config", postgresql15Controller.SaveConfig)
			route.Get("userConfig", postgresql15Controller.GetUserConfig)
			route.Post("userConfig", postgresql15Controller.SaveUserConfig)
			route.Get("log", postgresql15Controller.Log)
			route.Post("clearLog", postgresql15Controller.ClearLog)
			route.Get("databases", postgresql15Controller.DatabaseList)
			route.Post("databases", postgresql15Controller.AddDatabase)
			route.Delete("databases", postgresql15Controller.DeleteDatabase)
			route.Get("backups", postgresql15Controller.BackupList)
			route.Post("backups", postgresql15Controller.CreateBackup)
			route.Put("backups", postgresql15Controller.UploadBackup)
			route.Delete("backups", postgresql15Controller.DeleteBackup)
			route.Post("backups/restore", postgresql15Controller.RestoreBackup)
			route.Get("users", postgresql15Controller.UserList)
			route.Post("users", postgresql15Controller.AddUser)
			route.Delete("users", postgresql15Controller.DeleteUser)
			route.Post("users/password", postgresql15Controller.SetUserPassword)
		})
		r.Prefix("postgresql16").Group(func(route route.Router) {
			postgresql16Controller := plugins.NewPostgresql16Controller()
			route.Get("status", postgresql16Controller.Status)
			route.Post("reload", postgresql16Controller.Reload)
			route.Post("start", postgresql16Controller.Start)
			route.Post("stop", postgresql16Controller.Stop)
			route.Post("restart", postgresql16Controller.Restart)
			route.Get("load", postgresql16Controller.Load)
			route.Get("config", postgresql16Controller.GetConfig)
			route.Post("config", postgresql16Controller.SaveConfig)
			route.Get("userConfig", postgresql16Controller.GetUserConfig)
			route.Post("userConfig", postgresql16Controller.SaveUserConfig)
			route.Get("log", postgresql16Controller.Log)
			route.Post("clearLog", postgresql16Controller.ClearLog)
			route.Get("databases", postgresql16Controller.DatabaseList)
			route.Post("databases", postgresql16Controller.AddDatabase)
			route.Delete("databases", postgresql16Controller.DeleteDatabase)
			route.Get("backups", postgresql16Controller.BackupList)
			route.Post("backups", postgresql16Controller.CreateBackup)
			route.Put("backups", postgresql16Controller.UploadBackup)
			route.Delete("backups", postgresql16Controller.DeleteBackup)
			route.Post("backups/restore", postgresql16Controller.RestoreBackup)
			route.Get("users", postgresql16Controller.UserList)
			route.Post("users", postgresql16Controller.AddUser)
			route.Delete("users", postgresql16Controller.DeleteUser)
			route.Post("users/password", postgresql16Controller.SetUserPassword)
		})
		r.Prefix("php74").Group(func(route route.Router) {
			php74Controller := plugins.NewPHPController(74)
			route.Get("status", php74Controller.Status)
			route.Post("reload", php74Controller.Reload)
			route.Post("start", php74Controller.Start)
			route.Post("stop", php74Controller.Stop)
			route.Post("restart", php74Controller.Restart)
			route.Get("load", php74Controller.Load)
			route.Get("config", php74Controller.GetConfig)
			route.Post("config", php74Controller.SaveConfig)
			route.Get("fpmConfig", php74Controller.GetFPMConfig)
			route.Post("fpmConfig", php74Controller.SaveFPMConfig)
			route.Get("errorLog", php74Controller.ErrorLog)
			route.Get("slowLog", php74Controller.SlowLog)
			route.Post("clearErrorLog", php74Controller.ClearErrorLog)
			route.Post("clearSlowLog", php74Controller.ClearSlowLog)
			route.Get("extensions", php74Controller.GetExtensionList)
			route.Post("extensions", php74Controller.InstallExtension)
			route.Delete("extensions", php74Controller.UninstallExtension)
		})
		r.Prefix("php80").Group(func(route route.Router) {
			php80Controller := plugins.NewPHPController(80)
			route.Get("status", php80Controller.Status)
			route.Post("reload", php80Controller.Reload)
			route.Post("start", php80Controller.Start)
			route.Post("stop", php80Controller.Stop)
			route.Post("restart", php80Controller.Restart)
			route.Get("load", php80Controller.Load)
			route.Get("config", php80Controller.GetConfig)
			route.Post("config", php80Controller.SaveConfig)
			route.Get("fpmConfig", php80Controller.GetFPMConfig)
			route.Post("fpmConfig", php80Controller.SaveFPMConfig)
			route.Get("errorLog", php80Controller.ErrorLog)
			route.Get("slowLog", php80Controller.SlowLog)
			route.Post("clearErrorLog", php80Controller.ClearErrorLog)
			route.Post("clearSlowLog", php80Controller.ClearSlowLog)
			route.Get("extensions", php80Controller.GetExtensionList)
			route.Post("extensions", php80Controller.InstallExtension)
			route.Delete("extensions", php80Controller.UninstallExtension)
		})
		r.Prefix("php81").Group(func(route route.Router) {
			php81Controller := plugins.NewPHPController(81)
			route.Get("status", php81Controller.Status)
			route.Post("reload", php81Controller.Reload)
			route.Post("start", php81Controller.Start)
			route.Post("stop", php81Controller.Stop)
			route.Post("restart", php81Controller.Restart)
			route.Get("load", php81Controller.Load)
			route.Get("config", php81Controller.GetConfig)
			route.Post("config", php81Controller.SaveConfig)
			route.Get("fpmConfig", php81Controller.GetFPMConfig)
			route.Post("fpmConfig", php81Controller.SaveFPMConfig)
			route.Get("errorLog", php81Controller.ErrorLog)
			route.Get("slowLog", php81Controller.SlowLog)
			route.Post("clearErrorLog", php81Controller.ClearErrorLog)
			route.Post("clearSlowLog", php81Controller.ClearSlowLog)
			route.Get("extensions", php81Controller.GetExtensionList)
			route.Post("extensions", php81Controller.InstallExtension)
			route.Delete("extensions", php81Controller.UninstallExtension)
		})
		r.Prefix("php82").Group(func(route route.Router) {
			php82Controller := plugins.NewPHPController(82)
			route.Get("status", php82Controller.Status)
			route.Post("reload", php82Controller.Reload)
			route.Post("start", php82Controller.Start)
			route.Post("stop", php82Controller.Stop)
			route.Post("restart", php82Controller.Restart)
			route.Get("load", php82Controller.Load)
			route.Get("config", php82Controller.GetConfig)
			route.Post("config", php82Controller.SaveConfig)
			route.Get("fpmConfig", php82Controller.GetFPMConfig)
			route.Post("fpmConfig", php82Controller.SaveFPMConfig)
			route.Get("errorLog", php82Controller.ErrorLog)
			route.Get("slowLog", php82Controller.SlowLog)
			route.Post("clearErrorLog", php82Controller.ClearErrorLog)
			route.Post("clearSlowLog", php82Controller.ClearSlowLog)
			route.Get("extensions", php82Controller.GetExtensionList)
			route.Post("extensions", php82Controller.InstallExtension)
			route.Delete("extensions", php82Controller.UninstallExtension)
		})
		r.Prefix("php83").Group(func(route route.Router) {
			php83Controller := plugins.NewPHPController(83)
			route.Get("status", php83Controller.Status)
			route.Post("reload", php83Controller.Reload)
			route.Post("start", php83Controller.Start)
			route.Post("stop", php83Controller.Stop)
			route.Post("restart", php83Controller.Restart)
			route.Get("load", php83Controller.Load)
			route.Get("config", php83Controller.GetConfig)
			route.Post("config", php83Controller.SaveConfig)
			route.Get("fpmConfig", php83Controller.GetFPMConfig)
			route.Post("fpmConfig", php83Controller.SaveFPMConfig)
			route.Get("errorLog", php83Controller.ErrorLog)
			route.Get("slowLog", php83Controller.SlowLog)
			route.Post("clearErrorLog", php83Controller.ClearErrorLog)
			route.Post("clearSlowLog", php83Controller.ClearSlowLog)
			route.Get("extensions", php83Controller.GetExtensionList)
			route.Post("extensions", php83Controller.InstallExtension)
			route.Delete("extensions", php83Controller.UninstallExtension)
		})
		r.Prefix("phpmyadmin").Group(func(route route.Router) {
			phpMyAdminController := plugins.NewPhpMyAdminController()
			route.Get("info", phpMyAdminController.Info)
			route.Post("port", phpMyAdminController.SetPort)
		})
		r.Prefix("pureftpd").Group(func(route route.Router) {
			pureFtpdController := plugins.NewPureFtpdController()
			route.Get("status", pureFtpdController.Status)
			route.Post("start", pureFtpdController.Start)
			route.Post("stop", pureFtpdController.Stop)
			route.Post("restart", pureFtpdController.Restart)
			route.Get("list", pureFtpdController.List)
			route.Post("add", pureFtpdController.Add)
			route.Delete("delete", pureFtpdController.Delete)
			route.Post("changePassword", pureFtpdController.ChangePassword)
			route.Get("port", pureFtpdController.GetPort)
			route.Post("port", pureFtpdController.SetPort)
		})
		r.Prefix("redis").Group(func(route route.Router) {
			redisController := plugins.NewRedisController()
			route.Get("status", redisController.Status)
			route.Post("start", redisController.Start)
			route.Post("stop", redisController.Stop)
			route.Post("restart", redisController.Restart)
			route.Get("load", redisController.Load)
			route.Get("config", redisController.GetConfig)
			route.Post("config", redisController.SaveConfig)
		})
		r.Prefix("s3fs").Group(func(route route.Router) {
			s3fsController := plugins.NewS3fsController()
			route.Get("list", s3fsController.List)
			route.Post("add", s3fsController.Add)
			route.Post("delete", s3fsController.Delete)
		})
		r.Prefix("supervisor").Group(func(route route.Router) {
			supervisorController := plugins.NewSupervisorController()
			route.Get("status", supervisorController.Status)
			route.Post("start", supervisorController.Start)
			route.Post("stop", supervisorController.Stop)
			route.Post("restart", supervisorController.Restart)
			route.Post("reload", supervisorController.Reload)
			route.Get("log", supervisorController.Log)
			route.Post("clearLog", supervisorController.ClearLog)
			route.Get("config", supervisorController.Config)
			route.Post("config", supervisorController.SaveConfig)
			route.Get("processes", supervisorController.Processes)
			route.Post("startProcess", supervisorController.StartProcess)
			route.Post("stopProcess", supervisorController.StopProcess)
			route.Post("restartProcess", supervisorController.RestartProcess)
			route.Get("processLog", supervisorController.ProcessLog)
			route.Post("clearProcessLog", supervisorController.ClearProcessLog)
			route.Get("processConfig", supervisorController.ProcessConfig)
			route.Post("processConfig", supervisorController.SaveProcessConfig)
			route.Post("deleteProcess", supervisorController.DeleteProcess)
			route.Post("addProcess", supervisorController.AddProcess)

		})
		r.Prefix("fail2ban").Group(func(route route.Router) {
			fail2banController := plugins.NewFail2banController()
			route.Get("status", fail2banController.Status)
			route.Post("start", fail2banController.Start)
			route.Post("stop", fail2banController.Stop)
			route.Post("restart", fail2banController.Restart)
			route.Post("reload", fail2banController.Reload)
			route.Get("jails", fail2banController.List)
			route.Post("jails", fail2banController.Add)
			route.Delete("jails", fail2banController.Delete)
			route.Get("jails/{name}", fail2banController.BanList)
			route.Post("unban", fail2banController.Unban)
			route.Post("whiteList", fail2banController.SetWhiteList)
			route.Get("whiteList", fail2banController.GetWhiteList)
		})
		r.Prefix("podman").Group(func(route route.Router) {
			controller := plugins.NewPodmanController()
			route.Get("status", controller.Status)
			route.Get("isEnabled", controller.IsEnabled)
			route.Post("enable", controller.Enable)
			route.Post("disable", controller.Disable)
			route.Post("start", controller.Start)
			route.Post("stop", controller.Stop)
			route.Post("restart", controller.Restart)
			route.Get("registryConfig", controller.GetRegistryConfig)
			route.Post("registryConfig", controller.UpdateRegistryConfig)
			route.Get("storageConfig", controller.GetStorageConfig)
			route.Post("storageConfig", controller.UpdateStorageConfig)
		})
		r.Prefix("rsync").Group(func(route route.Router) {
			rsyncController := plugins.NewRsyncController()
			route.Get("status", rsyncController.Status)
			route.Post("start", rsyncController.Start)
			route.Post("stop", rsyncController.Stop)
			route.Post("restart", rsyncController.Restart)
			route.Get("modules", rsyncController.List)
			route.Post("modules", rsyncController.Create)
			route.Post("modules/{name}", rsyncController.Update)
			route.Delete("modules/{name}", rsyncController.Destroy)
			route.Get("config", rsyncController.GetConfig)
			route.Post("config", rsyncController.UpdateConfig)
		})
		r.Prefix("frp").Group(func(route route.Router) {
			frpController := plugins.NewFrpController()
			route.Get("status", frpController.Status)
			route.Get("isEnabled", frpController.IsEnabled)
			route.Post("enable", frpController.Enable)
			route.Post("disable", frpController.Disable)
			route.Post("start", frpController.Start)
			route.Post("stop", frpController.Stop)
			route.Post("restart", frpController.Restart)
			route.Get("config", frpController.GetConfig)
			route.Post("config", frpController.UpdateConfig)
		})
		r.Prefix("gitea").Group(func(route route.Router) {
			giteaController := plugins.NewGiteaController()
			route.Get("status", giteaController.Status)
			route.Get("isEnabled", giteaController.IsEnabled)
			route.Post("enable", giteaController.Enable)
			route.Post("disable", giteaController.Disable)
			route.Post("start", giteaController.Start)
			route.Post("stop", giteaController.Stop)
			route.Post("restart", giteaController.Restart)
			route.Get("config", giteaController.GetConfig)
			route.Post("config", giteaController.UpdateConfig)
		})
		r.Prefix("toolbox").Group(func(route route.Router) {
			toolboxController := plugins.NewToolBoxController()
			route.Get("dns", toolboxController.GetDNS)
			route.Post("dns", toolboxController.SetDNS)
			route.Get("swap", toolboxController.GetSWAP)
			route.Post("swap", toolboxController.SetSWAP)
			route.Get("timezone", toolboxController.GetTimezone)
			route.Post("timezone", toolboxController.SetTimezone)
			route.Get("hosts", toolboxController.GetHosts)
			route.Post("hosts", toolboxController.SetHosts)
			route.Post("rootPassword", toolboxController.SetRootPassword)
		})
	})
}
