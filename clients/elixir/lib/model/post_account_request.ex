# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule .Model.PostAccountRequest do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :accountId,
    :name
  ]

  @type t :: %__MODULE__{
    :accountId => String.t,
    :name => String.t | nil
  }
end

defimpl Poison.Decoder, for: .Model.PostAccountRequest do
  def decode(value, _options) do
    value
  end
end

