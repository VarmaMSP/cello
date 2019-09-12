import React, { Component } from 'react'
import ButtonWithIcon from '../../button_with_icon'

export interface Props {
  handleClose: () => void
  children: JSX.Element | JSX.Element[]
}

export default class Modal extends Component<Props> {
  componentDidMount() {
    const scrollY = document.documentElement.style.getPropertyValue(
      '--scroll-y',
    )

    document.body.style.position = 'fixed'
    document.body.style.top = `-${scrollY}`
  }

  componentWillUnmount() {
    const scrollY = document.body.style.top

    document.body.style.position = ''
    document.body.style.top = ''
    window.scrollTo(0, parseInt(scrollY || '0') * -1)
  }

  render() {
    const { children, handleClose } = this.props

    return (
      <>
        <div
          className="fixed inset-0 bg-white z-10"
          style={{ opacity: 0.99 }}
        />
        <div className="modal px-3 py-4 border md:border-2 md:border-gray-300 border-gray-400 shadow z-20">
          <div className="w-full h-5 relative">
            <ButtonWithIcon
              className="absolute right-0 w-4 text-gray-600"
              icon="close"
              onClick={handleClose}
            />
          </div>
          <div className="h-full">{children}</div>
        </div>
      </>
    )
  }
}
