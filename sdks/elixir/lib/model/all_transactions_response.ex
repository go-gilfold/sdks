# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule .Model.AllTransactionsResponse do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :transactions
  ]

  @type t :: %__MODULE__{
    :transactions => [.Model.TransactionResponse.t]
  }
end

defimpl Poison.Decoder, for: .Model.AllTransactionsResponse do
  import .Deserializer
  def decode(value, options) do
    value
    |> deserialize(:transactions, :list, .Model.TransactionResponse, options)
  end
end

