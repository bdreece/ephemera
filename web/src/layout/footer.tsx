export default function Footer() {
    return (
        <footer class="footer sm:footer-horizontal footer-center bg-base-300 text-base-content p-4">
            <aside>
                <p>
                    Copyright &copy; {new Date().getFullYear()} - All right
                    reserved by ACME Industries Ltd
                </p>
            </aside>
        </footer>
    );
}
