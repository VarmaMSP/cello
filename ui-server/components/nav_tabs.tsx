import classNames from 'classnames'
import React from 'react'

interface OwnProps {
  tabs: string[]
  active: string
}

const NavTabs: React.FC<OwnProps> = ({ tabs, active }) => {
  return (
    <div className="flex mt-8 mb-4 ">
      {tabs.map((t) => (
        <div
          className={classNames('mr-4 px-3 py-1 text-sm rounded-full', {
            'bg-green-300': t === active,
            'bg-gray-200': t !== active,
          })}
        >
          {t}
        </div>
      ))}
    </div>
  )
}

export default NavTabs
