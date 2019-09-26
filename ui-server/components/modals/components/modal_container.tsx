import ButtonWithIcon from 'components/button_with_icon'
import React, { useEffect, useRef } from 'react'

interface Props {
  handleClose: () => void
  closeUponClicking: 'OVERLAY' | 'CROSS'
  children: JSX.Element | JSX.Element[]
}

const ModalContainer: React.SFC<Props> = (props) => {
  const ref = useRef(null) as React.RefObject<HTMLDivElement>
  const { handleClose, closeUponClicking } = props

  const handleClickOutside = (e: any) => {
    if (ref.current && !ref.current.contains(e.target as Node)) {
      closeUponClicking === 'OVERLAY' && handleClose()
    }
  }

  useEffect(() => {
    document.addEventListener('mousedown', handleClickOutside)
    return () => {
      document.removeEventListener('mousedown', handleClickOutside)
    }
  })

  return (
    <div
      ref={ref}
      className="modal px-4 py-4 bg-white border md:border-2 md:border-gray-300 border-gray-400 shadow z-20"
    >
      {closeUponClicking === 'CROSS' && (
        <div className="w-full h-5 relative">
          <ButtonWithIcon
            className="absolute right-0 w-4 text-gray-600 hover:text-black"
            icon="close"
            onClick={handleClose}
          />
        </div>
      )}
      <div className="h-full">{props.children}</div>
    </div>
  )
}

export default ModalContainer
