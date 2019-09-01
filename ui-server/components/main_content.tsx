import React from 'react'

interface Props {
  children: JSX.Element | JSX.Element[]
}

const MainContent: React.SFC<Props> = ({ children }) => (
  <div className="lg:pl-56 pl-4 lg:pr-5 pr-4 pt-20 pb-64 z-0">
    <div className="lg:pl-5 lg:pb-36">{children}</div>
  </div>
)

export default MainContent
