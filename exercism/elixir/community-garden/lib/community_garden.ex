# Use the Plot struct as it is provided
defmodule Plot do
  @enforce_keys [:plot_id, :registered_to]
  defstruct [:plot_id, :registered_to]
end

defmodule CommunityGarden do
  def start() do
    Agent.start(fn -> {0, []} end)
  end

  def list_registrations(pid) do
    Agent.get(pid, fn {_, plots} -> plots end)
  end

  def register(pid, register_to) do
    Agent.get_and_update(pid, fn {count, plots} ->
      next_id = count + 1
      plot = %Plot{plot_id: next_id, registered_to: register_to}
      {plot, {next_id, [plot | plots]}}
    end)
  end

  def release(pid, plot_id) do
    Agent.update(pid, fn {count, plots} ->
      new_state = Enum.filter(plots, &(&1.plot_id != plot_id))

      {count, new_state}
    end)
  end

  def get_registration(pid, plot_id) do
    Agent.get(pid, fn {_count, plots} ->
      plots
      |> Enum.filter(&(&1.plot_id == plot_id))
      |> List.first()
      |> maybe_return_not_found()
    end)
  end

  defp maybe_return_not_found(nil), do: {:not_found, "plot is unregistered"}

  defp maybe_return_not_found(%Plot{} = plot), do: plot
end
