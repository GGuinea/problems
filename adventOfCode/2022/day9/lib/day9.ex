defmodule Day9 do
  @moduledoc false

  def get_input do
    {:ok, body} = File.read("input")
    body |> String.split("\n", trim: true)
  end

  def part_1 do
    get_input()
    |> Enum.flat_map(fn row ->
      [dir, count] = String.split(row, " ")
      List.duplicate(String.to_atom(dir), String.to_integer(count))
    end)
    |> simulate([{{0,0} {0,0}}])
  end

  defp simulate(input, position) do
  end
end
