import { ChevronDownIcon } from '@heroicons/react/outline'
import { displaySymbol } from '@utils/displaySymbol'

import {
  getMenuItemHoverBgForCoin,
  getBorderStyleForCoinHover,
} from '@styles/coins'

export default function SelectTokenDropdown({ chainId, selected, onClick }) {
  const symbol = displaySymbol(chainId, selected)

  return (
    <button
      className="sm:mt-[-1px] flex-shrink-0 mr-[-1px] w-[35%] cursor-pointer focus:outline-none"
      onClick={onClick}
    >
      <div
        className={`
          group rounded-xl
          -ml-2
          bg-white bg-opacity-10
        `}
      >
        <div
          className={`
            flex justify-center md:justify-start 
            bg-[#49444c] bg-opacity-100
            transform-gpu transition-all duration-100
            ${getMenuItemHoverBgForCoin(selected)}
            border border-transparent
            ${getBorderStyleForCoinHover(selected)}
            items-center 
            rounded-lg
            py-1.5 pl-2 h-14
          `}
        >
          <div className="self-center flex-shrink-0 hidden mr-1 sm:block">
            <div className="relative flex p-1 rounded-full">
              <img className="rounded-md w-7 h-7" src={selected.icon} />
            </div>
          </div>
          <div className="text-left cursor-pointer">
            <h4 className="text-lg font-medium text-white">
              <span>{symbol}</span>
              <ChevronDownIcon className="inline w-4 ml-2 -mt-1 transition-all transform focus:rotate-180" />
            </h4>
          </div>
        </div>
      </div>
    </button>
  )
}