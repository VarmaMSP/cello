import { Link } from 'components/link'
import React from 'react'
import Categories from './categories'
import Recommended from './recommended'

const HomeView: React.FC<{}> = () => {
  return (
    <div>
      <Recommended />
      <Categories />
      <div className="md:hidden text-center">
        <hr className="my-5" />
        <p className="leading-tight">
          <Link href="/about" prefetch={false}>
            <a className="cursor-pointer">{'about'}</a>
          </Link>{' '}
          <span className="font-extrabold">&middot;</span>{' '}
          <Link href="/privacy" prefetch={false}>
            <a className="cursor-pointer">{'privacy'}</a>
          </Link>
        </p>
        <p className="mb-1">
          <a href="https://www.facebook.com/phenopod" target="_blank">
            {'facebook'}
          </a>{' '}
          <span className="font-extrabold">&middot;</span>{' '}
          <a href="https://twitter.com/phenopod" target="_blank">
            {'twitter'}
          </a>{' '}
          <span className="font-extrabold">&middot;</span>{' '}
          <a href="https://www.reddit.com/r/phenopod/" target="_blank">
            {'reddit'}
          </a>
        </p>
        <a href="mailto:hello@phenopod.com" className="font-light">
          {'hello@phenopod.com'}
        </a>
      </div>
    </div>
  )
}

export default HomeView
