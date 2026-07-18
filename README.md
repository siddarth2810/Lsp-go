# LSP


## Connecting to nvim
```lua
                local client = vim.lsp.start_client({
                        cmd = { "/home/sid/personal/lsp/main" },
                        name = "sidlsp",
                        on_attach = require("sid.lsp").on_attach,
                })

                if not client then
                        vim.notify("hey you did not do the sid lsp thing")
                        return
                end

                vim.api.nvim_create_autocmd("FileType", {
                        pattern = "markdown",
                        callback = function ()
                                vim.lsp.buf_attach_client(0, client)
                        end,
                })

```
