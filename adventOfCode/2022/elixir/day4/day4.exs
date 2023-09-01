defmodule Day4 do
  def get_input() do
    {:ok, body} = File.read("input.prod")

    body
    |> String.split("\n", trim: true)
    |> Enum.map(&String.split(&1, ","))
    |> Enum.map(&conv/1)
  end

  def part_1(),
    do:
      get_input()
      |> Enum.map(&is_overlap/1)
      |> Enum.sum()

  def part_2(),
    do:
      get_input()
      |> Enum.map(&contain/1)
      |> Enum.sum()

  defp is_overlap(params) do
    if(
      (params.start1 <= params.start2 and params.end1 >= params.end2) or
        (params.start2 <= params.start1 and params.end2 >= params.end1)
    ) do
      1
    else
      0
    end
  end

  defp conv([head | tail]) do
    [s1, e1] = String.split(head, "-")
    [s2, e2] = String.split(hd(tail), "-")

    %{
      start1: String.to_integer(s1),
      end1: String.to_integer(e1),
      start2: String.to_integer(s2),
      end2: String.to_integer(e2)
    }
  end

  defp contain(params) do
    if params.start1 in params.start2..params.end2 or params.end1 in params.start2..params.end2 or
         params.start2 in params.start1..params.end1 or params.end2 in params.start1..params.end1 do
      1
    else
      0
    end
  end
end
