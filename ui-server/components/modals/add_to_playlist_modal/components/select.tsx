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
    <>
      {options.map((o) => (
        <div
          className="flex items-center cursor-pointer"
          onClick={() => handleSelect(o.id)}
        >
          <div className="w-4 h-4 mx-3 rounded-full border-2 border-blue-600">
            {o.id === selected && (
              <div className="w-full h-full bg-blue-400 rounded-full" />
            )}
          </div>
          {o.display}
        </div>
      ))}
    </>
  )
}

export default Select
