import classNames from 'classnames'
import React from 'react'

interface OwnProps {
  tabs: string[]
  active: string
  onClick: (tab: string) => void
}

const NavTabs: React.FC<OwnProps> = ({ tabs, active, onClick }) => {
  return (
    <div className="flex mt-8 mb-4 ">
      {tabs.map((t) => (
        <div
          className={classNames('mr-4 px-3 py-1 text-sm rounded-full', {
            'bg-green-300': t === active,
            'bg-gray-200': t !== active,
            'cursor-default': t === active,
            'cursor-pointer': t !== active,
          })}
          onClick={() => onClick(t)}
        >
          {t}
        </div>
      ))}
    </div>
  )
}

export default NavTabs
