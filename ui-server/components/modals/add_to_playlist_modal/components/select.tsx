import React from 'react'

interface Props {
  options: {
    id: string
    display: string
  }[]
  selected: string
  handleSelect: (id: string) => void
}

const Select: React.FC<Props> = ({ options, selected, handleSelect }) => {
  return (
    <div>
      {options.map((o) => (
        <div key={o.id}>
          <label className="inline-flex items-center">
            <input
              type="checkbox"
              className="form-checkbox"
              checked={o.id === selected}
              onClick={() => handleSelect(o.id)}
            />
            <span className="ml-2">{o.display}</span>
          </label>
        </div>
      ))}
    </div>
  )
}

export default Select
  