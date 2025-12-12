"use client";

export default function Spaceman() {
    return (
        <div className="grid grid-cols-5 grid-rows-6 gap-2 border border-emerald-700 p-7 mx-auto rounded-xs container mb-15 text-gray-900 bg-purple-100">
            <div className="col-span-3 row-span-6 p-2">
                <span className="bg-purple-500 rounded-xs p-1 text-xs">Spaces</span>
                <div></div>
            </div>
            <div className="col-span-2 row-span-6 col-start-4">
                <span className="bg-purple-500 rounded-xs p-1 text-xs">
                    Trending discussions
                </span>
            </div>
        </div>
    );
}
