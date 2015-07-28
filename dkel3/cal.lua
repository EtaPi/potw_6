#! /usr/bin/env lua
-- cal.lua

-- Get arguments. Print usage and exit if passed incorrect arguments
year = arg[1]
width = arg[2]
if year == nil or width == nil then
  print("Usage: cal.lua <year> <width>")
  error()
end

-- Calculate the number of leap years that occurred between 0 and the input year
function num_leaps(year)
  local num_mod4 = math.floor( (year - 1)/4 )
  local num_mod100 = math.floor( (year - 1)/100 )
  local num_mod400 = math.floor( (year - 1)/400 )

  local leaps = num_mod4 - num_mod100 + num_mod400
  return leaps
end

-- Calculate the starting day offset for the input year
function year_offset(year)
  return (year + num_leaps(year)) % 7
end

-- Returns true if the input year is a leap year
function is_leap_year( year )
  local is_zero_mod4 = (year % 4 == 0)
  local is_zero_mod100 = (year % 100 == 0)
  local is_zero_mod400 = (year % 400 == 0)
  local is_leap_year = (is_zero_mod4 and not is_zero_mod100) or is_zero_mod400

  return is_leap_year
end

-- Given an input string and a width, pads the string with blank spaces to
-- match the input width.
function center(str, width)
  local len = str:len()
  local padding_left = math.floor( (width - len)/2 )
  local padding_right = width - padding_left - len

  local ret = string.rep(" ", padding_left) .. str .. string.rep(" ", padding_right)

  return ret
end

-- Give a month name, the starting day offset, and the number of days in the
-- month, pretty prints the month and returns it as an array of strings
function gen_month (name, offset, num_days)
  local month = {}

  -- Month name header
  table.insert( month, center(name, 20) )
  -- Day of the week header
  table.insert( month, "Su Mo Tu We Th Fr Sa" )

  local start = 1 - offset

  -- six rows
  for row=1,6 do
    local row_string = ""

    -- seven columns
    for day=1,7 do
      -- calculate the number of the day for sunday row 1 (could be negative)
      local this_day = start + (row - 1) * 7 + (day - 1)

      if (this_day > 0 and this_day <= num_days) then
        -- add day number if its valid
        row_string = row_string .. string.format("%2d", this_day)
      else
        -- otherwise add blank spaces
        row_string = row_string .. "  "
      end

      -- add padding space between columns
      if not (day == 7) then
        row_string = row_string .. " "
      end
    end

    table.insert( month, row_string )
  end

  return month
end

-- Given an input year and width, prints a calendar with 'width' number of
-- columns
function print_year( year, width )
  local total_width = width * 20 + 3 * (width + 1)
  header = center( tostring(year), total_width)
  print(header)

  local offset = year_offset(year)

  -- twelve months
  for i=1,12 do
    local row = {}

    -- loop to concatenate months into same row
    for w=1,width do
      local index = (i - 1) * width + (w - 1) + 1
      if index > 12 then break end

      local month = month_days[index]
      local month_string = gen_month( month['name'], offset, month['days'] )

      -- increment offset for next month
      offset = ( offset + month['days'] ) % 7

      row = merge_months( row, month_string )
    end

    -- advance month counter based on width
    i = i + width

    -- print everything
    for m in pairs(row) do
      print("   " .. row[m] .. "   ")
    end
  end
end

-- Given two arrays representing pretty printed months, merges them into one
-- array by concatenating each line.
function merge_months( m1, m2)
  local ret = {}

  for row=1,8 do
    local new_row = ""

    -- concatenate operator doesnt like nils
    if (m1[row] == nil) then
      new_row = m2[row]
    elseif (m2[row] == nil) then
      new_row = m1[row]
    else
      new_row = m1[row] .. "   " .. m2[row]
    end

    table.insert( ret, new_row )
  end

  return ret
end

-- Set up an array with the number of days in each month
month_days = {}
table.insert( month_days, {name='January',    days=31})
table.insert( month_days, {name='February',   days=28})
table.insert( month_days, {name='March',      days=31})
table.insert( month_days, {name='April',      days=30})
table.insert( month_days, {name='May',        days=31})
table.insert( month_days, {name='June',       days=30})
table.insert( month_days, {name='July',       days=31})
table.insert( month_days, {name='August',     days=31})
table.insert( month_days, {name='September',  days=30})
table.insert( month_days, {name='October',    days=31})
table.insert( month_days, {name='November',   days=30})
table.insert( month_days, {name='December',   days=31})

-- If leap year, make February have 29 days
if (is_leap_year(year)) then
  month_days[2] = 29
end

-- do werk
print_year(year, width)
