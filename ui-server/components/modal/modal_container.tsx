import classNames from 'classnames'
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
  header?: string
  closeUponClicking?: 'OVERLAY' | 'CROSS'
  className: string
  children: JSX.Element | JSX.Element[]
}

const ModalContainer: React.SFC<DispatchToProps & OwnProps> = ({
  closeModal,
  header,
  closeUponClicking,
  className,
  children,
}) => {
  const ref = useRef(null) as React.RefObject<HTMLDivElement>

  const handleClickOutside = (e: any) => {
    if (ref.current && !ref.current.contains(e.target as Node)) {
      ;(closeUponClicking === undefined || closeUponClicking === 'OVERLAY') &&
        closeModal()
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
      className={classNames(
        'flex flex-col bg-white border shadow z-20',
        className,
      )}
    >
      <div className="w-full flex flex-row-reverse flex-none items-center px-4 py-3">
        {(closeUponClicking === undefined || closeUponClicking === 'CROSS') && (
          <ButtonWithIcon
            className="flex-none w-4 text-gray-600 hover:text-black"
            icon="close"
            onClick={closeModal}
          />
        )}
        {header && (
          <h4 className="flex-1 text-2xl text-gray-900 tracking-wide">
            {header}
          </h4>
        )}
      </div>

      <div className="flex-1 px-4 py-3">{children}</div>
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
