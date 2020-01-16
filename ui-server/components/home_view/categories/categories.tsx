import { ChartLink } from 'components/link'
import React from 'react'
import { Curation } from 'types/app'

export interface StateToProps {
  categories: Curation[]
}

const Categories: React.FC<StateToProps> = ({ categories }) => {
  return (
    <div>
      <h2 className="text-xl text-gray-700">{'Browse by categories'}</h2>
      <hr className="mt-2 mb-4" />
      <ul>
        {categories.map((c) => (
          <ChartLink key={c.id} chartId={c.id}>
            <a>
              <li className="block flex items-center h-10 lg:w-1/2 pl-6 rounded-full hover:bg-gray-200">
                <span className="text-lg text-gray-900">{c.title}</span>
                &nbsp;&nbsp;
                <span className="md:inline hidden text-sm text-gray-700 tracking-wider">
                  {c.subTitle}
                </span>
              </li>
            </a>
          </ChartLink>
        ))}
      </ul>
    </div>
  )
}

export default Categories
