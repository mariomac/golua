-- Says "hello" followed by the string passed as argument
-- @param name The string to follow the "hello" message
-- @return The number of characters of the 'name' string
function sayhello(name)
    host_print('hello ' .. name) -- This function has to be provided by the host language
    return string.len(name)
end


