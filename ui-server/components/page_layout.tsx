import React from 'react'

interface OwnProps {
  children: [JSX.Element, JSX.Element]
}

const PageLayout: React.FC<OwnProps> = ({ children }) => {
  return (
    <div className="flex md:flex-row flex-col">
      <div className="lg:w-2/3 w-full lg:mr-5">{children[0]}</div>
      <div className="lg:w-1/3">{children[1]}</div>
    </div>
  )
}

export default PageLayout
