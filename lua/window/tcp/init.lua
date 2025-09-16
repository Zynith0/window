local M = {}

M.create_window = function(input)
	local max_height = vim.api.nvim_win_get_height(0)
	local max_width = vim.api.nvim_win_get_width(0)

	local height = math.floor(max_height * input)
	local width = math.floor(max_width * input)

	local buf = vim.api.nvim_create_buf(false, false)

	vim.api.nvim_open_win(buf, true, {
		relative = "editor",
		height = height,
		width = width,
		col = (max_width - width) / 2,
		row = (max_height - height) / 2,
		border = "double",
	})
end

vim.api.nvim_create_user_command("NewWin", "lua require('tcp').newWin()", {})

M.newWin = function()
	local command_output = vim.fn.system("nc localhost 8080")
	local cmd = vim.api.nvim_replace_termcodes(command_output, true, false, true)
	vim.api.nvim_feedkeys(cmd, "n", false)
end

return M
