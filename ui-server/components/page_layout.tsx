import React from 'react'

interface OwnProps {
  children: JSX.Element | [JSX.Element, JSX.Element]
}

export default class PageLayout extends React.Component<OwnProps> {
  componentDidCatch() {
    !!window && window.location.reload(true)
  }

  render() {
    const { children } = this.props

    return Array.isArray(children) ? (
      <div className="page-layout-split">
        <div className="first">{children[0]}</div>
        <div className="second">{children[1]}</div>
      </div>
    ) : (
      <div className="page-layout">{children}</div>
    )
  }
}
