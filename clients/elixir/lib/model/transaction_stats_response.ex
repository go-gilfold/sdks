# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule .Model.TransactionStatsResponse do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :total,
    :daily,
    :monthly
  ]

  @type t :: %__MODULE__{
    :total => float(),
    :daily => [.Model.TransactionStatsResponseDailyInner.t],
    :monthly => [.Model.TransactionStatsResponseDailyInner.t]
  }
end

defimpl Poison.Decoder, for: .Model.TransactionStatsResponse do
  import .Deserializer
  def decode(value, options) do
    value
    |> deserialize(:daily, :list, .Model.TransactionStatsResponseDailyInner, options)
    |> deserialize(:monthly, :list, .Model.TransactionStatsResponseDailyInner, options)
  end
end

