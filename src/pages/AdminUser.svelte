<script lang="ts">
    import '../app.css'
    import '../js/declares'
    import AdminNavbar from '../components/AdminNavbar.svelte'
    import Footer from '../components/Footer.svelte'
    import DataTable from '../components/DataTable.svelte'
    import { onMount } from 'svelte'

    export let siteName: string
    
    let newUsername = '';
    let errorMessage = '';
    let successMessage = '';
    let dataTableInstance: any;
    let isDataTablesLoaded = false;

    // 模擬假資料
    let users = [
        {
            id: "1",
            username: "admin",
            created_at: "2024-03-15T08:00:00Z"
        },
        {
            id: "2",
            username: "editor01",
            created_at: "2024-03-16T10:30:00Z"
        },
        {
            id: "3",
            username: "writer_alice",
            created_at: "2024-03-17T14:15:00Z"
        },
        {
            id: "4",
            username: "moderator_bob",
            created_at: "2024-03-18T09:45:00Z"
        },
        {
            id: "5",
            username: "tester_carol",
            created_at: "2024-03-19T16:20:00Z"
        }
    ];

    const tableId = 'users-table';

    // 修改 tableConfig
    $: tableConfig = {
        data: users,
        columns: [
            { 
                data: 'id', 
                title: 'ID' 
            },
            { 
                data: 'username', 
                title: '用戶名' 
            },
            { 
                data: 'created_at', 
                title: '建立時間',
                render: (data) => new Date(data).toLocaleString()
            },
            {
                data: null,
                title: '操作',
                render: (data) => `<button class="text-red-500 hover:text-red-700 delete-user" data-id="${data.id}">刪除</button>`
            }
        ],
        searching: true,
        ordering: true,
        paging: true,
        language: {
            "lengthMenu": "顯示 _MENU_ 條記錄",
            "zeroRecords": "沒有找到記錄",
            "info": "第 _PAGE_ 頁 ( 總共 _PAGES_ 頁 )",
            "infoEmpty": "無記錄",
            "infoFiltered": "(從 _MAX_ 條記錄過濾)",
            "search": "搜尋:",
            "paginate": {
                "first": "首頁",
                "last": "末頁",
                "next": "下一頁",
                "previous": "上一頁"
            }
        }
    };

    async function addUser() {
        if (!newUsername.trim()) {
            errorMessage = '請輸入用戶名';
            return;
        }

        const newUser = {
            id: (users.length + 1).toString(),
            username: newUsername,
            created_at: new Date().toISOString()
        };

        users = [...users, newUser];
        
        // 如果 DataTable 實例存在，更新表格
        if (dataTableInstance) {
            dataTableInstance.clear().rows.add(users).draw();
        }

        successMessage = '用戶新增成功';
        newUsername = '';
        
        setTimeout(() => {
            successMessage = '';
        }, 3000);
    }

    async function deleteUser(userId: string) {
        if (!confirm('確定要刪除此用戶嗎？')) return;

        users = users.filter(user => user.id !== userId);
        
        // 如果 DataTable 實例存在，更新表格
        if (dataTableInstance) {
            dataTableInstance.clear().rows.add(users).draw();
        }

        successMessage = '用戶刪除成功';
        
        setTimeout(() => {
            successMessage = '';
        }, 3000);
    }
</script>

<div class="min-h-screen flex flex-col">
    <AdminNavbar {siteName} />
    
    <main class="flex-grow container mx-auto px-4 py-8">
        <h1 class="text-2xl font-bold mb-6">用戶管理</h1>

        <!-- 新增用戶表單 -->
        <div class="bg-white p-6 rounded-lg shadow-md mb-8">
            <h2 class="text-xl font-semibold mb-4">新增用戶</h2>
            <div class="flex gap-4">
                <input
                    type="text"
                    bind:value={newUsername}
                    placeholder="輸入用戶名"
                    class="flex-grow p-2 border rounded"
                />
                <button
                    on:click={addUser}
                    class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
                >
                    新增
                </button>
            </div>
            
            {#if errorMessage}
                <p class="text-red-500 mt-2">{errorMessage}</p>
            {/if}
            {#if successMessage}
                <p class="text-green-500 mt-2">{successMessage}</p>
            {/if}
        </div>

        <!-- 用戶列表 -->
        <div class="bg-white p-6 rounded-lg shadow-md">
            <h2 class="text-xl font-semibold mb-4">用戶列表</h2>
            <table id={tableId} class="display">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>用戶名</th>
                        <th>建立時間</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>
                    {#each users as user}
                    <tr>
                        <td>{user.id}</td>
                        <td>{user.username}</td>
                        <td>{new Date(user.created_at).toLocaleString()}</td>
                        <td>
                            <button class="text-red-500 hover:text-red-700 delete-user" data-id={user.id}>
                                刪除
                            </button>
                        </td>
                    </tr>
                    {/each}
                </tbody>
            </table>
            <DataTable id={tableId} config={tableConfig} />
        </div>
    </main>

    <Footer {siteName} />
</div>
