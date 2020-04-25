package config

import "time"

/**
 * log配置信息
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func log() map[string]interface{} {
	return map[string]interface{}{
		"dir":                env("app_log.dir", nil),
		"filename":           env("app_log.filename", nil),
		"service_filename":   env("app_log.service_filename", nil),
		"conductor_filename": env("app_log.conductor_filename", nil),
		"backup_filename":    env("app_log.backup_filename", nil),
		"max_age":            7 * 24 * time.Hour,
		"rotation_time":      24 * time.Hour,
	}
}
