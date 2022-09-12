defmodule Recursion do
  def tuple_to_number(number) do
    number |> Tuple.to_list() |> List.to_integer(2)
  end

  def o2([number], _pos) do
    tuple_to_number(number)
  end

  def o2(numbers, pos) do
    zero_count = Enum.count(numbers, &(elem(&1, pos) == ?0))
    one_count = length(numbers) - zero_count
    to_keep = if one_count >= zero_count, do: ?1, else: ?0
    numbers = Enum.filter(numbers, &(elem(&1, pos) == to_keep))
    o2(numbers, pos + 1)
  end

  def co2([number], _pos) do
    tuple_to_number(number)
  end

  def co2(numbers, pos) do
    zero_count = Enum.count(numbers, &(elem(&1, pos) == ?0))
    one_count = length(numbers) - zero_count
    to_keep = if one_count >= zero_count, do: ?0, else: ?1
    numbers = Enum.filter(numbers, &(elem(&1, pos) == to_keep))
    co2(numbers, pos + 1)
  end
end

defmodule Day3 do
  import Bitwise

  def part_1() do
    input = get_input()
    [sample | _] = input
    digit_length = tuple_size(sample)
    half = div(length(input), 2)

    number_as_list =
      for i <- 0..(digit_length - 1) do
        zero_count = Enum.count_until(input, &(elem(&1, i) == ?0), half + 1)
        if zero_count > half, do: ?0, else: ?1
      end

    gamma = List.to_integer(number_as_list, 2)

    mask = 2 ** digit_length - 1

    epsilon =
      ~~~gamma &&&
        mask

    gamma * epsilon
  end

  def part_2() do
    o2 = Recursion.o2(get_input(), 0)
    co2 = Recursion.co2(get_input(), 0)
    IO.inspect(o2 * co2)
  end

  def get_input() do
    {:ok, body} = File.read("input")

    body
    |> String.split("\n", trim: true)
    |> Enum.map(&(&1 |> String.to_charlist() |> List.to_tuple()))
  end
end

IO.inspect(Day3.part_1())
IO.inspect(Day3.part_2())
