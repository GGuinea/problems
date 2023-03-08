defmodule TakeANumber do
  def start() do
    spawn(&loop/0)
  end

  defp loop(state \\ 0) do
    receive do
      {:report_state, sender} ->
        send(sender, state)
        loop(state)

      {:take_a_number, sender} ->
        state = state + 1
        send(sender, state)
        loop(state)

      :stop ->
        :ok

      _ ->
        loop(state)
    end
  end
end
