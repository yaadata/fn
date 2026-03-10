local function add_optionsgo_dot_import()
	local bufnr = vim.api.nvim_get_current_buf()
	if vim.bo[bufnr].filetype ~= "go" then
		return
	end

	local lines = vim.api.nvim_buf_get_lines(bufnr, 0, -1, false)
	local target = '. "github.com/yaadata/optionsgo"'

	for _, line in ipairs(lines) do
		if line:find("github.com/yaadata/optionsgo", 1, true) then
			return
		end
	end

	local import_start, import_end
	for i, line in ipairs(lines) do
		if line:match("^import%s*%($") then
			import_start = i
		elseif import_start and line:match("^%)$") then
			import_end = i
			break
		end
	end

	if import_start and import_end then
		vim.api.nvim_buf_set_lines(bufnr, import_end - 1, import_end - 1, false, {
			"\t" .. target,
		})
		return
	end

	for i, line in ipairs(lines) do
		if line:match("^package%s+") then
			vim.api.nvim_buf_set_lines(bufnr, i, i, false, {
				"",
				"import " .. target,
			})
			return
		end
	end
end

vim.api.nvim_create_user_command("GolangAddOptionsGoDotImport", add_optionsgo_dot_import, {})
