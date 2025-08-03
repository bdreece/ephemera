export default function Sidebar() {
    return (
        <>
            <span class="sidebar-overlay"></span>

            <aside class="drawer left">
                <div class="content flex flex-col h-full">
                    <button class="btn sm circle ghost absolute right-2 top-2">
                        &times;
                    </button>
                    <h2 class="text-xl">Sidebar</h2>
                </div>
            </aside>
        </>
    );
}
