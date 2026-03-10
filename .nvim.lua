local root = vim.fs.dirname(debug.getinfo(1, "S").source:sub(2))
local config_dir = root .. "/tools/neovim"

local function load_project_configs()
  local paths = vim.fn.globpath(config_dir, "*.lua", false, true)
  -- Load shared project config first, then let *.local.lua files override it.
  table.sort(paths, function(a, b)
    local a_is_local = a:match("%.local%.lua$") ~= nil
    local b_is_local = b:match("%.local%.lua$") ~= nil

    if a_is_local ~= b_is_local then
      return not a_is_local
    end

    return a < b
  end)

  for _, path in ipairs(paths) do
    dofile(path)
  end
end

if vim.uv.fs_stat(config_dir) then
  load_project_configs()
end
