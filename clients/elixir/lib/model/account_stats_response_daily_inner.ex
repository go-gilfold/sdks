# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule .Model.AccountStatsResponseDailyInner do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :date,
    :amount
  ]

  @type t :: %__MODULE__{
    :date => String.t,
    :amount => float()
  }
end

defimpl Poison.Decoder, for: .Model.AccountStatsResponseDailyInner do
  def decode(value, _options) do
    value
  end
end

