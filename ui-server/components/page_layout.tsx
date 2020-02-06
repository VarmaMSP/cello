import React from 'react'

interface OwnProps {
  children: JSX.Element | [JSX.Element, JSX.Element]
}

const PageLayout: React.FC<OwnProps> = ({
  children,
}) => {
  return Array.isArray(children) ? (
    <div className="page-layout-split">
      <div className="first">{children[0]}</div>
      <div className="second">{children[1]}</div>
    </div>
  ) : (
    <div className="page-layout">{children}</div>
  )
}

export default PageLayout
