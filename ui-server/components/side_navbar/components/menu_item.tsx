import classnames from 'classnames'
import { Icon, iconMap } from 'components/icon'
import { Link } from 'components/link'
import { useRouter } from 'next/router'
import React from 'react'

interface Props {
  icon: Icon
  name: string
  href: string
}

const MenuItem: React.SFC<Props> = ({ icon, name, href }) => {
  const Icon = iconMap[icon]
  const currentUrlPath = useRouter().asPath

  return (
    <Link href={href} scroll={false}>
      <a
        className={classnames(
          'flex items-center w-full h-full px-4 my-2 tracking-wide cursor-pointer',
          {
            'text-teal-800': href !== currentUrlPath,
            'text-red-600 bg-red-100 rounded-full': href === currentUrlPath,
          },
        )}
      >
        <Icon className="w-4 h-4 mr-3 fill-current" />
        <h4 className="capitalize text-lg font-medium leading-loose tracking-wide">
          {name}
        </h4>
      </a>
    </Link>
  )
}

export default MenuItem
