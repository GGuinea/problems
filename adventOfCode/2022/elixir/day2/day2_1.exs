defmodule Day2 do
  def get_input() do
    {:ok, body} = File.read("input")

    body
    |> String.split("\n", trim: true)
    |> Enum.map(&String.split(&1, " "))
  end

  def part1(),
    do:
      get_input()
      |> Enum.map(&match_part_1/1)
      |> Enum.sum()

  def part2(),
    do:
      get_input()
      |> Enum.map(&match_part_2/1)
      |> Enum.sum()

  defp match_part_1(to_match) do
    # A - rock
    # B - paper
    # C - scissor

    # X - rock 1
    # Y - paper 2
    # Z - scissors 3

    # X - lose
    # Y - draw
    # Z -  win
    case to_match do
      # A - rock
      ["A", "X"] -> 4
      ["A", "Y"] -> 8
      ["A", "Z"] -> 3
      ["B", "X"] -> 1
      ["B", "Y"] -> 5
      ["B", "Z"] -> 9
      ["C", "X"] -> 7
      ["C", "Y"] -> 2
      ["C", "Z"] -> 6
    end
  end

  defp match_part_2(to_match) do
    # A - rock
    # B - paper
    # C - scissor

    # X - rock 1
    # Y - paper 2
    # Z - scissors 3

    # X - lose
    # Y - draw
    # Z -  win

    case to_match do
      ["A", "X"] -> 3
      ["A", "Y"] -> 4
      ["A", "Z"] -> 8
      ["B", "X"] -> 1
      ["B", "Y"] -> 5
      ["B", "Z"] -> 9
      ["C", "X"] -> 2
      ["C", "Y"] -> 6
      ["C", "Z"] -> 7
    end
  end
end
