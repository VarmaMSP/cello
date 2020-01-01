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
  active: string | undefined
  defaultTab?: string
}

const NavTabs: React.FC<OwnProps> = ({ tabs, active, defaultTab }) => {
  return (
    <div className="flex border-b">
      {tabs.map((t) => (
        <div key={t.name} className="w-20 mr-2 text-center">
          <Link
            href={{ pathname: t.pathname, query: t.query }}
            as={t.as}
            key={t.name}
          >
            <a
              className={classNames(
                'block px-3 py-1 text-sm capitalize leading-loose',
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
              'bg-green-500':
                (active !== undefined && t.name === active) ||
                (active === undefined && t.name === defaultTab),
            })}
          />
        </div>
      ))}
    </div>
  )
}

export default NavTabs
