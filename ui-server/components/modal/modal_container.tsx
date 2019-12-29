import { StateToProps } from 'components/add_to_playlist_modal/add_to_playlist_modal'
import ButtonWithIcon from 'components/button_with_icon'
import React, { Dispatch, useEffect, useRef } from 'react'
import { connect } from 'react-redux'
import { AppState } from 'store'
import { AppActions, MODAL_MANAGER_CLOSE_MODAL } from 'types/actions'

interface DispatchToProps {
  closeModal: () => void
}

interface OwnProps {
  closeUponClicking: 'OVERLAY' | 'CROSS'
  children: JSX.Element | JSX.Element[]
}

const ModalContainer: React.SFC<DispatchToProps & OwnProps> = ({
  closeModal,
  closeUponClicking,
  children,
}) => {
  const ref = useRef(null) as React.RefObject<HTMLDivElement>

  const handleClickOutside = (e: any) => {
    if (ref.current && !ref.current.contains(e.target as Node)) {
      closeUponClicking === 'OVERLAY' && closeModal()
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
      className="modal md:px-6 px-4 py-6 bg-white border shadow z-20"
    >
      {closeUponClicking === 'CROSS' && (
        <div className="w-full h-5 relative">
          <ButtonWithIcon
            className="absolute right-0 w-4 text-gray-600 hover:text-black"
            icon="close"
            onClick={closeModal}
          />
        </div>
      )}
      <div className="h-full">{children}</div>
    </div>
  )
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    closeModal: () => dispatch({ type: MODAL_MANAGER_CLOSE_MODAL }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  null,
  mapDispatchToProps,
)(ModalContainer)
