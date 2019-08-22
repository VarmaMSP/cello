import React from 'react'

interface Props {
  children: JSX.Element | JSX.Element[]
}

const MainContent: React.SFC<Props> = ({ children }) => (
  <div className="lg:pl-56 pt-16 pb-32 z-0">{children}</div>
)

export default MainContent
