defmodule Day7 do
  @moduledoc false
  @total 70_000_000
  @need 30_000_000

  def get_input() do
    {:ok, body} = File.read("input.prod")

    body
    |> String.split("\n", trim: true)
  end

  def part_1 do
    get_input()
    |> Enum.map(&String.split(&1, " "))
    |> build_tree(%{}, [])
    |> Map.values()
    |> Enum.filter(fn x -> x < 100_000 end)
    |> Enum.sum()
  end

  def part_2 do
    values =
      get_input()
      |> Enum.map(&String.split(&1, " "))
      |> build_tree(%{}, [])
      |> Map.values()
      |> Enum.sort(:desc)

    free = @total - hd(values)
    still_need = @need - free
    values |> Enum.filter(fn x -> x >= still_need end) |> Enum.sort(:asc) |> hd()
  end

  def build_tree([], tree, _path), do: tree

  def build_tree([head | tail], tree, path) do
    case head do
      ["$", "cd", ".."] ->
        {_elem, path} = List.pop_at(path, -1)
        build_tree(tail, tree, path)

      ["$", "cd", "/"] ->
        build_tree(tail, tree, path ++ ["/"])

      ["$", "cd", dir_name] ->
        path = path ++ ["/" <> dir_name]

        build_tree(tail, tree, path)

      ["$", "ls"] ->
        build_tree(tail, tree, path)

      ["dir", _dir_name] ->
        build_tree(tail, tree, path)

      [size, _file_name] ->
        paths_to_update = get_subpaths_to_update(path)

        tree = paths_to_update |> Enum.reduce(tree, fn x, acc -> update_maps(acc, size, x) end)

        build_tree(tail, tree, path)
    end
  end

  defp update_maps(acc, size, path),
    do:
      acc
      |> insert_key_if_not_exist(path)
      |> Map.update(path, 0, fn existing_value ->
        existing_value + String.to_integer(size)
      end)

  defp get_subpaths_to_update(path),
    do:
      1..length(path)
      |> Enum.map(&retreiv_one_subpath(path, &1))
      |> Enum.map(&List.to_string/1)

  defp retreiv_one_subpath(path, x) do
    {res, _} = Enum.split(path, x)
    res
  end

  defp insert_key_if_not_exist(tree, key) do
    if not Map.has_key?(tree, key) do
      Map.put(tree, key, 0)
    else
      tree
    end
  end
end
