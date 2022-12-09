defmodule Day8 do
  @moduledoc false

  def get_input() do
    {:ok, body} = File.read("input.prod")
    body |> String.split("\n", trim: true)
  end

  def part_1 do
    list_of_lists =
      get_input()
      |> Enum.map(&String.graphemes/1)

    forest = get_tuple(list_of_lists)

    rows = Kernel.tuple_size(forest)
    cols = Kernel.tuple_size(elem(forest, 0))

    rows_seq = 0..(rows - 1)
    cols_seq = 0..(cols - 1)

    transponed_list = List.zip(list_of_lists) |> List.to_tuple()

    # forest
    # |> Enum.reduce(%{}, fn x, acc -> build_grid(x, acc) end)
    build_gird(rows_seq, cols_seq, forest)
    |> List.flatten()
    |> Enum.reduce(%{}, fn x, acc -> Map.merge(x, acc) end)
    |> Enum.reduce(0, fn x, acc ->
      acc + find_trees(x, rows_seq, cols_seq, acc, forest, transponed_list)
    end)
  end

  defp find_trees(entry, rows_size, cols_size, _acc, rows, cols) do
    {{x, y}, height} = entry

    if is_edge?(x, y, rows_size, cols_size) do
      1
    else
      is_seen(String.to_integer(height), elem(rows, x), elem(cols, y), x, y)
    end
  end

  defp is_seen(height, col, row, x, y) do
    {left, right} = col |> Tuple.to_list() |> Enum.split(y)
    {up, down} = row |> Tuple.to_list() |> Enum.split(x)

    if check_left(left, height) or check_left(tl(right), height) or
         check_left(up, height) or check_left(tl(down), height) do
      1
    else
      0
    end
  end

  defp check_left(list, height) do
    list = Enum.map(list, &String.to_integer/1)

    if Enum.all?(list, fn x -> x < height end) do
      # IO.inspect("found  #{list} heid: #{height}")
      true
    else
      # IO.inspect("Not seen in: #{list} heid: #{height}")
      false
    end
  end

  defp get_int_if_needed(x) when is_integer(x), do: x

  defp get_int_if_needed(x) when is_binary(x), do: String.to_integer(x)

  defp is_edge?(x, y, rows_size, cols_size) do
    _first..last_row = rows_size
    _first..last_col = cols_size
    # IO.inspect("x#{x} y#{y} rowsize#{last_row} colsize#{last_col}")

    if x == 0 or y == 0 or x == last_row or y == last_col do
      true
    else
      false
    end
  end

  def part_2 do
    list_of_lists =
      get_input()
      |> Enum.map(&String.graphemes/1)

    forest = get_tuple(list_of_lists)

    rows = Kernel.tuple_size(forest)
    cols = Kernel.tuple_size(elem(forest, 0))

    rows_seq = 0..(rows - 1)
    cols_seq = 0..(cols - 1)

    transponed_list = List.zip(list_of_lists) |> List.to_tuple()

    res =
      build_gird(rows_seq, cols_seq, forest)
      |> List.flatten()
      |> Enum.reduce(%{}, fn x, acc -> Map.merge(x, acc) end)
      |> Enum.reduce([], fn x, acc ->
        acc ++ [find_trees_2(x, rows_seq, cols_seq, acc, forest, transponed_list)]
      end)

    # IO.inspect(res)
    res |> Enum.max()
  end

  defp find_trees_2(entry, rows_size, cols_size, _acc, rows, cols) do
    {{x, y}, height} = entry

    if is_edge?(x, y, rows_size, cols_size) do
      0
    else
      is_seen_2(String.to_integer(height), elem(rows, x), elem(cols, y), x, y)
    end
  end

  defp is_seen_2(height, col, row, x, y) do
    {left, right} = col |> Tuple.to_list() |> Enum.split(y)
    {up, down} = row |> Tuple.to_list() |> Enum.split(x)

    check_left_2(Enum.reverse(left), height, y) * check_left_2(tl(right), height, y) *
      check_left_2(Enum.reverse(up), height, x) * check_left_2(tl(down), height, x)
  end

  defp check_left_2(list, height, index) do
    res =
      list
      |> Enum.map(&String.to_integer/1)
      |> Enum.find_index(fn x -> x >= height end)

    # |> Enum.reduce(0, fn x, acc -> acc + can_see(x, height) end)
    res = map_nil(res, length(list))

    IO.inspect("list: #{list}")
    IO.inspect("height: #{height}")
    IO.inspect("res: #{res}")
    IO.inspect("index: #{index}")
    res
  end

  defp map_nil(x, len) when is_nil(x) do
    len
  end

  defp map_nil(x, _len) when is_integer(x) do
    if x == 0 do
      1
    else
      x + 1
    end
  end

  defp can_see(x, height) do
    if height >= x do
      1
    else
      0
    end
  end

  defp get_tuple(list_of_lists) do
    list_of_lists
    |> Enum.map(&List.to_tuple/1)
    |> List.to_tuple()
  end

  defp build_gird(rows, cols, input) do
    grid = %{}

    for x <- rows do
      for y <- cols do
        Map.put(grid, {x, y}, input |> elem(x) |> elem(y))
      end
    end
  end
end
