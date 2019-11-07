import classNames from 'classnames'
import Link from 'next/link'
import React from 'react'

interface Tab {
  name: string
  pathname: string
  query: { [x: string]: number | string | boolean }
  as: string
}

interface OwnProps {
  tabs: Tab[]
  active: string
}

const NavTabs: React.FC<OwnProps> = ({ tabs, active }) => {
  return (
    <div className="flex mt-8 mb-4 ">
      {tabs.map((t) => (
        <Link href={{ pathname: t.pathname, query: t.query }} as={t.as}>
          <a
            className={classNames('block mr-4 px-3 py-1 text-sm rounded-full', {
              'bg-green-300': t.name === active,
              'bg-gray-200': t.name !== active,
              'cursor-default': t.name === active,
              'cursor-pointer': t.name !== active,
            })}
          >
            {t.name}
          </a>
        </Link>
      ))}
    </div>
  )
}

export default NavTabs
