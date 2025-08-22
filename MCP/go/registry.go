package main

import (
	"github.com/proxy-api/mcp-server/config"
	"github.com/proxy-api/mcp-server/models"
	tools_execute "github.com/proxy-api/mcp-server/tools/execute"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_execute.CreateDeleteproxyTool(cfg),
		tools_execute.CreateGetproxyTool(cfg),
		tools_execute.CreateOptionsproxyTool(cfg),
		tools_execute.CreatePatchproxyTool(cfg),
		tools_execute.CreatePostproxyTool(cfg),
		tools_execute.CreatePutproxyTool(cfg),
	}
}
