import { ChartLink } from 'components/link'
import React from 'react'
import { Chart } from 'types/app'

export interface StateToProps {
  categories: Chart[]
}

const Categories: React.FC<StateToProps> = ({ categories }) => {
  return (
    <div>
      <h2 className="text-xl text-gray-700">{'Browse by categories'}</h2>
      <hr className="mt-2 mb-4 border-gray-400" />
      <ul>
        {categories.map((c) => (
          <ChartLink chartId={c.id}>
            <a>
              <li
                key={c.id}
                className="block flex items-center h-10 md:w-1/2 pl-6 rounded-full hover:bg-gray-200"
              >
                <span className="text-lg text-gray-900">{c.title}</span>
                &nbsp;&nbsp;
                <span className="text-sm text-gray-700 tracking-wider">
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
