defmodule Day6 do
  @moduledoc false
  def get_input() do
    {:ok, body} = File.read("input.prod")

    body
    |> String.trim()
  end

  def part_1() do
    get_input()
    |> String.graphemes()
    |> Enum.chunk_every(4, 1)
    |> Enum.with_index()
    |> Enum.map(&find_first_unique(&1, 4))
  end

  def part_2() do
    get_input()
    |> String.graphemes()
    |> Enum.chunk_every(14, 1)
    |> Enum.with_index()
    |> Enum.map(&find_first_unique(&1, 14))
  end

  defp find_first_unique(str, chunk_by) do
    {string, index} = str
    res = Enum.uniq(string) |> length()
    if res == chunk_by, do: raise(Integer.to_string(index + chunk_by))
  end
end
