defmodule Day2 do
  def part_1() do
    {horizonatal, depth} =
      get_input()
      |> Enum.reduce({0, 0}, fn command, {horizonatal, depth} ->
        case command do
          {:forward, steps} -> {horizonatal + steps, depth}
          {:down, steps} -> {horizonatal, depth + steps}
          {:up, steps} -> {horizonatal, depth - steps}
        end
      end)

    IO.inspect(horizonatal * depth)
  end

  def part_2() do
    {horizonatal, depth, aim} =
      get_input()
      |> Enum.reduce({0, 0, 0}, fn command, {horizonatal, depth, aim} ->
        case command do
          {:forward, steps} -> {horizonatal + steps, depth + aim * steps, aim}
          {:down, steps} -> {horizonatal, depth, aim + steps}
          {:up, steps} -> {horizonatal, depth, aim - steps}
        end
      end)

    IO.inspect(horizonatal * depth)
  end

  def get_input() do
    {:ok, body} = File.read("input")

    body
    |> String.split("\n", trim: true)
    |> Enum.map(fn line ->
      [direction, steps] = String.split(line)
      {String.to_atom(direction), String.to_integer(steps)}
    end)
  end
end

Day2.part_1()
Day2.part_2()
