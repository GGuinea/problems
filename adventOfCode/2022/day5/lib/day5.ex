defmodule Day5 do
  @test_start %{
    "1" => ["Z", "N"],
    "2" => ["M", "C", "D"],
    "3" => ["P"]
  }

  @prod_start %{
    "1" => ["H", "B", "V", "W", "N", "M", "L", "P"],
    "2" => ["M", "Q", "H"],
    "3" => ["N", "D", "B", "G", "F", "Q", "M", "L"],
    "4" => ["Z", "T", "F", "Q", "M", "W", "G"],
    "5" => ["M", "T", "H", "P"],
    "6" => ["C", "B", "M", "J", "D", "H", "G", "T"],
    "7" => ["M", "N", "B", "F", "V", "R"],
    "8" => ["P", "L", "H", "M", "R", "G", "S"],
    "9" => ["P", "D", "B", "C", "N"]
  }
  def get_input() do
    {:ok, body} = File.read("input.prod")

    body
    |> String.split("\n", trim: true)
    |> Enum.map(&String.split(&1, ","))
    |> List.flatten()
  end

  def part_1() do
    get_input()
    |> prepare_data
    |> Enum.reduce(@prod_start, fn x, acc -> make_move(x, acc, true) end)
    |> concat_last_values()
  end

  def part_2() do
    get_input()
    |> prepare_data
    |> Enum.reduce(@prod_start, fn x, acc -> make_move(x, acc, false) end)
    |> concat_last_values()
  end

  defp prepare_data(input),
    do:
      input
      |> Enum.map(&Regex.scan(~r{\d+}, &1))
      |> IO.inspect()
      |> Enum.map(&create_map/1)

  defp concat_last_values(input),
    do:
      input
      |> Map.values()
      |> Enum.map(&List.last/1)
      |> Enum.reduce("", fn x, acc -> acc <> x end)

  defp make_move(x, acc, reverse) do
    {head, tail} = Enum.split(acc[x.from], -String.to_integer(x.number))
    acc = Kernel.put_in(acc[x.from], head)

    if reverse do
      Kernel.put_in(acc[x.to], acc[x.to] ++ Enum.reverse(tail))
    else
      Kernel.put_in(acc[x.to], acc[x.to] ++ tail)
    end
  end

  defp create_map([first, second, third]) do
    %{
      number: hd(first),
      from: hd(second),
      to: hd(third)
    }
  end
end
