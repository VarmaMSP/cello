import React, { cloneElement, useEffect, useState } from 'react'

interface Props {
  button: JSX.Element
  dropdown: JSX.Element | JSX.Element[]
}

const Dropdown: React.SFC<Props> = (props) => {
  const dropdownRef = React.createRef<HTMLDivElement>()
  const [showDropDown, setShowDropdown] = useState(false)
  // TODO: Hack to close dropdown when cliked on button. Try to find a better way
  const [skipOnClick, setSkipOnClick] = useState(false)

  useEffect(() => {
    document.addEventListener('mousedown', (e) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(e.target as any)
      ) {
        setShowDropdown(false)
        setSkipOnClick(true)
      }
    })
  })

  const { button, dropdown } = props

  return (
    <div className="relative w-full h-full">
      {cloneElement(button, {
        onClick: () => {
          if (skipOnClick) {
            setSkipOnClick(false)
            return
          }
          setShowDropdown(!showDropDown)
        },
        className: `${button.props.className} absolute right-0`,
      })}
      {showDropDown && (
        <div
          className="absolute right-0 z-50 py-5 bg-white border border-gray-400 shadow rounded"
          style={{ top: '130%' }}
          ref={dropdownRef}
        >
          {dropdown}
        </div>
      )}
    </div>
  )
}

export default Dropdown
