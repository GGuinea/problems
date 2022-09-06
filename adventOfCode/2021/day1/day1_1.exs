defmodule Day1 do
  def part_1() do
    input =
      get_input()
      |> count_increase_records()
  end

  def part_2() do
    input =
      get_input()
      |> Enum.chunk_every(3, 1, :discard)
      |> Enum.map(&Enum.sum/1)
      |> count_increase_records()
  end

  def count_increase_records(list) do
    list
    |> Enum.chunk_every(2, 1, :discard)
    |> Enum.count(fn [a, b] -> b > a end)
    |> IO.inspect()
  end

  def get_input() do
    {:ok, body} = File.read("input")

    body
    |> String.split("\n", trim: true)
    |> Enum.map(&String.to_integer/1)
  end
end

Day1.part_1()
Day1.part_2()
