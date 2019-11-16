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
    <div className="flex">
      {tabs.map((t) => (
        <div className="w-20 mr-2 text-center">
          <Link
            href={{ pathname: t.pathname, query: t.query }}
            as={t.as}
            key={t.name}
          >
            <a
              className={classNames(
                'block px-3 py-1 text-sm capitalize rounded-br-lg rounded-bl-lg',
                {
                  'cursor-default': t.name === active,
                  'cursor-pointer': t.name !== active,
                },
              )}
            >
              {t.name}
            </a>
          </Link>
          <div
            className={classNames('h-1 w-20 rounded-full', {
              'bg-green-500': t.name === active,
            })}
          />
        </div>
      ))}
    </div>
  )
}

export default NavTabs
