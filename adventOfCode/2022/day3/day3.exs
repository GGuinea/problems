defmodule Day3 do
  def get_input() do
    {:ok, body} = File.read("input.prod")

    body
    |> String.split("\n", trim: true)
  end

  defp match(var) do
    if var > 96 do
      var - 96
    else
      var - 38
    end
  end

  def part_2(),
    do:
      get_input()
      |> Enum.chunk_every(3)
      |> Enum.map(&get_badge/1)
      |> Enum.map(&common_in_three/1)
      |> common_op

  defp common_op(arg),
    do:
      arg
      |> List.flatten()
      |> Enum.reduce(fn x, acc -> x <> acc end)
      |> String.to_charlist()
      |> Enum.map(&match/1)
      |> Enum.sum()

  defp get_badge(list),
    do:
      list
      |> Enum.map(&String.graphemes(&1))

  defp common_in_three([first, second, third] = list),
    do:
      first
      |> Enum.filter(fn el -> Enum.member?(second, el) end)
      |> Enum.filter(fn el -> Enum.member?(third, el) end)
      |> Enum.uniq()

  def part_1(),
    do:
      get_input()
      |> Enum.map(&String.graphemes(&1))
      |> Enum.map(&Enum.chunk_every(&1, div(length(&1), 2)))
      |> Enum.map(&List.myers_difference(Enum.uniq(hd(&1)), Enum.uniq(hd(tl(&1)))))
      |> Enum.map(&Keyword.get(&1, :eq))
      |> common_op
end
