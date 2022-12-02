defmodule Day1 do
  def get_input() do
    {:ok, body} = File.read("input")

    body
    |> String.split("\n\n", trim: true)
  end

  def part1(),
    do:
      sum_per_elf()
      |> Enum.max()

  def part2(),
    do:
      sum_per_elf()
      |> Enum.sort()
      |> Enum.take(-3)
      |> Enum.sum()

  def sum_per_elf(),
    do:
      get_input()
      |> Enum.map(&String.split(&1, "\n", trim: true))
      |> Enum.map(&Enum.map(&1, fn y -> String.to_integer(y) end))
      |> Enum.map(&Enum.reduce(&1, fn x, acc -> acc + x end))
end
